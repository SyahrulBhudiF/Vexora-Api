package user

import (
	"fmt"
	"github.com/SyahrulBhudiF/Vexora-Api/internal/domains/user/entity"
	userRepositories "github.com/SyahrulBhudiF/Vexora-Api/internal/domains/user/repository"
	"github.com/SyahrulBhudiF/Vexora-Api/internal/helpers"
	"github.com/SyahrulBhudiF/Vexora-Api/internal/services"
	"github.com/SyahrulBhudiF/Vexora-Api/internal/types"
	"github.com/SyahrulBhudiF/Vexora-Api/internal/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
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
