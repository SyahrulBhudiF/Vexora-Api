package services

import (
	"fmt"
	"github.com/SyahrulBhudiF/Vexora-Api/internal/domains/user/entity"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"time"
)

type UserClaims struct {
	UUID     uuid.UUID `json:"uuid"`
	Username string    `json:"username"`
	jwt.StandardClaims
}

type JWTService struct {
	secret string
}

func NewJWTService(secret string) *JWTService {
	return &JWTService{secret}
}

func (j *JWTService) GenerateAccessToken(user entity.User, expTime time.Duration) (string, error) {
	claims := UserClaims{
		UUID:     user.UUID,
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(expTime).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	jwt, err := token.SignedString([]byte(j.secret))

	if err != nil {
		return "", err
	}

	return jwt, nil
}

func (j *JWTService) GenerateRefreshToken(user entity.User, expTime time.Duration) (string, error) {
	claims := jwt.StandardClaims{
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: time.Now().Add(expTime).Unix(),
		Subject:   user.UUID.String(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	jwt, err := token.SignedString([]byte(j.secret))

	if err != nil {
		return "", err
	}

	return jwt, nil
}

func (j *JWTService) ValidateAccessToken(token string) (*UserClaims, error) {
	claims := &UserClaims{}

	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.secret), nil
	})

	if err != nil {
		return nil, err
	}

	// validate claims structure
	if claims.UUID == uuid.Nil || claims.Username == "" {
		return nil, fmt.Errorf("invalid access token")
	}

	// validate expiration
	if time.Now().Unix() > claims.ExpiresAt {
		return nil, fmt.Errorf("access token expired")
	}

	return claims, nil
}

func (j *JWTService) ValidateRefreshToken(token string) (*jwt.StandardClaims, error) {
	claims := &jwt.StandardClaims{}

	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.secret), nil
	})

	if err != nil {
		return nil, err
	}

	// validate claims structure
	if claims.Subject == "" {
		return nil, fmt.Errorf("invalid refresh token")
	}

	// validate expiration
	if time.Now().Unix() > claims.ExpiresAt {
		return nil, fmt.Errorf("refresh token expired")
	}

	return claims, nil
}

func (j *JWTService) Unserialize(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.secret), nil
	})
}
