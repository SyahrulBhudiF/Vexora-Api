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
	if exists := handler.userRepo.Exists(&user); exists {
		return helpers.ErrorResponse(ctx, fiber.StatusConflict, true, fmt.Errorf("username has been taken"))
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
