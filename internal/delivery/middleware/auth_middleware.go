package middleware

import (
	"fmt"
	"strings"

	"github.com/SyahrulBhudiF/Vexora-Api/internal/domains/user/entity"
	userRepositories "github.com/SyahrulBhudiF/Vexora-Api/internal/domains/user/repository"
	"github.com/SyahrulBhudiF/Vexora-Api/internal/helpers"
	"github.com/SyahrulBhudiF/Vexora-Api/internal/services"
	"github.com/SyahrulBhudiF/Vexora-Api/internal/types"
	"github.com/gofiber/fiber/v2"
)

type AuthMiddleware struct {
	userRepository  *userRepositories.UserRepository
	tokenRepository *types.RedisRepository
	jwtService      *services.JWTService
}

func NewAuthMiddleware(userRepository *userRepositories.UserRepository, tokenRepository *types.RedisRepository, jwtService *services.JWTService) *AuthMiddleware {
	return &AuthMiddleware{
		userRepository:  userRepository,
		tokenRepository: tokenRepository,
		jwtService:      jwtService,
	}
}

func (m *AuthMiddleware) EnsureAuthenticated(ctx *fiber.Ctx) error {
	var err error

	// parse request header
	authHeader := ctx.Get("Authorization", "")
	if authHeader == "" {
		return helpers.ErrorResponse(ctx, fiber.StatusUnauthorized, true, fmt.Errorf("authorization header is required"))
	}

	// parse bearer token
	authStr := strings.Split(authHeader, "Bearer ")

	if len(authStr) < 2 {
		return helpers.ErrorResponse(ctx, fiber.StatusUnauthorized, true, err)
	}

	token := authStr[1]

	// validate token
	claims, err := m.jwtService.ValidateAccessToken(token)
	if err != nil {
		return helpers.ErrorResponse(ctx, fiber.StatusUnauthorized, true, err)
	}

	user := entity.User{
		Entity: types.Entity{
			UUID: claims.UUID,
		},
	}

	if err := m.userRepository.Find(&user); err != nil {
		return helpers.ErrorResponse(ctx, fiber.StatusUnauthorized, true, err)
	}

	ctx.Locals("accessToken", token)
	ctx.Locals("user", &user)

	return ctx.Next()
}
