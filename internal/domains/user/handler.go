package user

import (
	"encoding/base64"
	"fmt"
	"github.com/SyahrulBhudiF/Vexora-Api/internal/domains/user/entity"
	userRepositories "github.com/SyahrulBhudiF/Vexora-Api/internal/domains/user/repository"
	"github.com/SyahrulBhudiF/Vexora-Api/internal/helpers"
	"github.com/SyahrulBhudiF/Vexora-Api/internal/services"
	"github.com/SyahrulBhudiF/Vexora-Api/internal/types"
	"github.com/SyahrulBhudiF/Vexora-Api/internal/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/spf13/viper"
	"io"
	"time"
)

type UserHandler struct {
	userRepo        *userRepositories.UserRepository
	tokenRepo       *types.RedisRepository
	jwtService      *services.JWTService
	imageKitService *services.ImageKitService
	viper           *viper.Viper
}

func NewUserHandler(userRepo *userRepositories.UserRepository, tokenRepo *types.RedisRepository, jwtService *services.JWTService, imageKitService *services.ImageKitService, viper *viper.Viper) *UserHandler {
	return &UserHandler{
		userRepo:        userRepo,
		tokenRepo:       tokenRepo,
		jwtService:      jwtService,
		imageKitService: imageKitService,
		viper:           viper,
	}
}

func (handler *UserHandler) Register(ctx *fiber.Ctx) error {
	body := ctx.Locals("body").(*RegisterRequest)

	user := entity.User{Username: body.Username}
	email := entity.User{Email: body.Email}
	if exists := handler.userRepo.Exists(&user) || handler.userRepo.Exists(&email); exists {
		return helpers.ErrorResponse(ctx, fiber.StatusConflict, true, fmt.Errorf("username or email has been taken"))
	}

	hashedPassword := utils.HashPassword(body.Password, handler.viper.GetString("app.secret"))

	user = *entity.NewUser(
		body.Username,
		body.Name,
		body.Email,
		hashedPassword,
		"",
		"",
	)

	if err := handler.userRepo.Create(&user); err != nil {
		return helpers.ErrorResponse(ctx, fiber.StatusConflict, true, fmt.Errorf("sign up failed"))
	}

	return ctx.JSON(types.WebResponse[entity.User]{Message: "sign up success!", Success: true, ShouldNotify: false, Data: user})
}

func (handler *UserHandler) Login(ctx *fiber.Ctx) error {
	body := ctx.Locals("body").(*LoginRequest)

	user := entity.User{Username: body.Username}
	if err := handler.userRepo.Find(&user); err != nil {
		return helpers.ErrorResponse(ctx, fiber.StatusUnauthorized, true, fmt.Errorf("invalid username or password"))
	}

	if !utils.ComparePassword(user.Password, body.Password, handler.viper.GetString("app.secret")) {
		return helpers.ErrorResponse(ctx, fiber.StatusUnauthorized, true, fmt.Errorf("invalid username or password"))
	}

	refreshTokenDuration := time.Duration(handler.viper.GetInt("auth.refresh_token_exp_days")) * time.Hour * 24
	refreshToken, err := handler.jwtService.GenerateRefreshToken(user, refreshTokenDuration)
	if err != nil {
		return helpers.ErrorResponse(ctx, fiber.StatusInternalServerError, true, fmt.Errorf("failed to generate refresh token"))
	}

	accessTokenDuration := time.Duration(handler.viper.GetInt("auth.access_token_exp_mins")) * time.Minute
	accessToken, err := handler.jwtService.GenerateAccessToken(user, accessTokenDuration)
	if err != nil {
		return helpers.ErrorResponse(ctx, fiber.StatusInternalServerError, true, fmt.Errorf("failed to generate access token"))
	}

	token := entity.Token{RefreshToken: refreshToken, AccessToken: accessToken}

	return ctx.JSON(types.WebResponse[entity.Token]{Message: "sign in success!", Success: true, ShouldNotify: false, Data: token})
}

func (handler *UserHandler) Logout(ctx *fiber.Ctx) error {
	body := ctx.Locals("body").(*LogoutRequest)
	rawAccessToken := ctx.Locals("accessToken").(string)

	refreshToken, err := handler.jwtService.ValidateRefreshToken(body.RefreshToken)
	if err != nil {
		return helpers.ErrorResponse(ctx, fiber.StatusUnauthorized, true, fmt.Errorf("invalid refresh token"))
	}

	accessToken, err := handler.jwtService.ValidateAccessToken(rawAccessToken)
	if err != nil {
		return helpers.ErrorResponse(ctx, fiber.StatusUnauthorized, true, fmt.Errorf("invalid access token"))
	}

	if refreshToken.Subject != accessToken.UUID.String() {
		return helpers.ErrorResponse(ctx, fiber.StatusUnauthorized, true, fmt.Errorf("permission denied"))
	}

	isBlacklisted, err := handler.tokenRepo.Exists(body.RefreshToken)
	if err != nil {
		return helpers.ErrorResponse(ctx, fiber.StatusInternalServerError, true, fmt.Errorf("failed to logout"))
	}

	if isBlacklisted {
		return helpers.ErrorResponse(ctx, fiber.StatusUnauthorized, true, fmt.Errorf("refresh token has been blacklisted"))
	}

	refreshTokenBlacklistDuration := time.Until(time.Unix(refreshToken.ExpiresAt, 0))
	if err := handler.tokenRepo.Set(body.RefreshToken, nil, refreshTokenBlacklistDuration); err != nil {
		return helpers.ErrorResponse(ctx, fiber.StatusInternalServerError, true, fmt.Errorf("failed to blacklist token"))
	}

	return helpers.SuccessResponse[any](ctx, fiber.StatusOK, false, "sign out success!", nil)
}

func (handler *UserHandler) GetProfile(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*entity.User)

	return helpers.SuccessResponse(ctx, fiber.StatusOK, false, "get profile success!", user)
}

func (handler *UserHandler) UpdateProfile(ctx *fiber.Ctx) error {
	body := ctx.Locals("body").(*UpdateProfileRequest)
	user := ctx.Locals("user").(*entity.User)

	if err := helpers.UpdateEntity(ctx, body, user); err != nil {
		return helpers.ErrorResponse(ctx, fiber.StatusInternalServerError, true, fmt.Errorf("failed to update profile"))
	}

	fmt.Println(user)

	if err := handler.userRepo.Update(user); err != nil {
		return helpers.ErrorResponse(ctx, fiber.StatusInternalServerError, true, fmt.Errorf("failed to update profile"))
	}

	return ctx.JSON(types.WebResponse[any]{Message: "update profile success!", Success: true, ShouldNotify: false})
}

func (handler *UserHandler) UploadProfilePicture(ctx *fiber.Ctx) error {
	file, err := ctx.FormFile("image")
	user := ctx.Locals("user").(*entity.User)

	if err != nil {
		return helpers.ErrorResponse(ctx, fiber.StatusBadRequest, true, fmt.Errorf("failed to get image"))
	}

	file.Filename = fmt.Sprintf("%s-%s", ctx.Locals("user").(*entity.User).UUID.String(), file.Filename)

	image, err := file.Open()

	if err != nil {
		return helpers.ErrorResponse(ctx, fiber.StatusInternalServerError, true, fmt.Errorf("failed to open image"))
	}

	defer image.Close()

	imageBuff, err := io.ReadAll(image)

	if err != nil {
		return helpers.ErrorResponse(ctx, fiber.StatusInternalServerError, true, fmt.Errorf("failed to read image"))
	}

	imageKitRes, err := handler.imageKitService.UploadImage(base64.StdEncoding.EncodeToString(imageBuff), "profiles", file.Filename)

	if err != nil {
		return helpers.ErrorResponse(ctx, fiber.StatusBadRequest, true, fmt.Errorf("failed to upload image, %v", err))
	}

	if user.FileId != "" {
		if err := handler.imageKitService.DeleteImage(user.FileId); err != nil {
			return helpers.ErrorResponse(ctx, fiber.StatusInternalServerError, true, fmt.Errorf("failed to delete old image"))
		}
	}

	user.ProfilePicture = imageKitRes.Data.Url
	user.FileId = imageKitRes.Data.FileId
	if err := handler.userRepo.Update(user); err != nil {
		return helpers.ErrorResponse(ctx, fiber.StatusInternalServerError, true, fmt.Errorf("failed to update profile"))
	}

	return helpers.SuccessResponse[any](ctx, fiber.StatusOK, false, "upload image success!", nil)
}

func (handler *UserHandler) ChangePassword(ctx *fiber.Ctx) error {
	body := ctx.Locals("body").(*ChangePasswordRequest)
	user := ctx.Locals("user").(*entity.User)

	if !utils.ComparePassword(user.Password, body.PreviousPassword, handler.viper.GetString("app.secret")) {
		return helpers.ErrorResponse(ctx, fiber.StatusUnauthorized, true, fmt.Errorf("invalid old password"))
	}

	hashedPassword := utils.HashPassword(body.NewPassword, handler.viper.GetString("app.secret"))
	user.Password = hashedPassword

	if err := handler.userRepo.Update(user); err != nil {
		return helpers.ErrorResponse(ctx, fiber.StatusInternalServerError, true, fmt.Errorf("failed to change password"))
	}

	return helpers.SuccessResponse[any](ctx, fiber.StatusOK, false, "change password success!", nil)
}

func (handler *UserHandler) RefreshToken(ctx *fiber.Ctx) error {
	body := ctx.Locals("body").(*RefreshTokenRequest)

	refreshClaims, err := handler.jwtService.ValidateRefreshToken(body.RefreshToken)
	if err != nil {
		return helpers.ErrorResponse(ctx, fiber.StatusUnauthorized, true, fmt.Errorf("invalid refresh token"))
	}

	isBlacklisted, err := handler.tokenRepo.Exists(body.RefreshToken)
	if err != nil {
		return helpers.ErrorResponse(ctx, fiber.StatusInternalServerError, true, fmt.Errorf("failed to check token blacklist"))
	}

	if isBlacklisted {
		return helpers.ErrorResponse(ctx, fiber.StatusUnauthorized, true, fmt.Errorf("token has been blacklisted"))
	}

	userUUID, err := uuid.Parse(refreshClaims.Subject)

	if err != nil {
		return helpers.ErrorResponse(ctx, fiber.StatusInternalServerError, true, fmt.Errorf("failed to parse user uuid"))
	}

	user := entity.User{
		Entity: types.Entity{UUID: userUUID},
	}

	if err := handler.userRepo.Find(&user); err != nil {
		return helpers.ErrorResponse(ctx, fiber.StatusUnauthorized, true, fmt.Errorf("user not found"))
	}

	accessTokenDuration := time.Duration(handler.viper.GetInt("auth.access_token_exp_mins")) * time.Minute
	accessToken, err := handler.jwtService.GenerateAccessToken(user, accessTokenDuration)

	if err != nil {
		return helpers.ErrorResponse(ctx, fiber.StatusInternalServerError, true, fmt.Errorf("failed to generate access token"))
	}

	token := entity.Token{RefreshToken: body.RefreshToken, AccessToken: accessToken}

	return helpers.SuccessResponse(ctx, fiber.StatusOK, false, "refresh token success!", token)
}
