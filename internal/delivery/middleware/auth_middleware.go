package middleware

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

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

	// check if user data is cached
	redisKey := fmt.Sprintf("user:%s", claims.UUID)
	cachedUser, err := m.tokenRepository.Get(redisKey)
	if err == nil && cachedUser != "" {
		var user entity.User
		if err := json.Unmarshal([]byte(cachedUser), &user); err == nil {
			ctx.Locals("accessToken", token)
			ctx.Locals("user", &user)
			return ctx.Next()
		}
	}

	// get user data from database if not cached
	user := entity.User{
		Entity: types.Entity{
			UUID: claims.UUID,
		},
	}

	if err := m.userRepository.Find(&user); err != nil {
		return helpers.ErrorResponse(ctx, fiber.StatusUnauthorized, true, err)
	}

	userJSON, err := json.Marshal(user)
	if err != nil {
		return helpers.ErrorResponse(ctx, fiber.StatusInternalServerError, true, fmt.Errorf("failed to cache user data"))
	}

	if err := m.tokenRepository.Set(redisKey, string(userJSON), 60*time.Minute); err != nil {
		return helpers.ErrorResponse(ctx, fiber.StatusInternalServerError, true, fmt.Errorf("failed to save user data to cache"))
	}

	ctx.Locals("accessToken", token)
	ctx.Locals("user", &user)
	return ctx.Next()
}
