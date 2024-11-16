package _interface

import (
	"github.com/SyahrulBhudiF/Vexora-Api/internal/domains/user/entity"
	"github.com/SyahrulBhudiF/Vexora-Api/internal/services"
	"github.com/golang-jwt/jwt"
	"time"
)

type IJWTService interface {
	// GenerateAccessToken generates a new access token for a user with specified expiration time
	GenerateAccessToken(user entity.User, expTime time.Duration) (string, error)

	// GenerateRefreshToken generates a new refresh token for a user with specified expiration time
	GenerateRefreshToken(user entity.User, expTime time.Duration) (string, error)

	// ValidateAccessToken validates an access token and returns the claims if valid
	ValidateAccessToken(token string) (*services.UserClaims, error)

	// ValidateRefreshToken validates a refresh token and returns the standard claims if valid
	ValidateRefreshToken(token string) (*jwt.StandardClaims, error)

	// Unserialize parses and validates a JWT token without checking specific claims
	Unserialize(token string) (*jwt.Token, error)
}
