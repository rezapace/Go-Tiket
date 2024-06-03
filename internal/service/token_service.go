package service

import (
	"Ticketing/common"
	"Ticketing/entity"
	"Ticketing/internal/config"
	"context"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenUsecase interface {
	GenerateAccessToken(ctx context.Context, user *entity.User) (string, error)
}

type TokenService struct {
	cfg *config.Config //ini dipake krna secret key nya diambil dari config
}

func NewTokenService(cfg *config.Config) *TokenService {
	return &TokenService{
		cfg: cfg,
	}
}

// untuk generate token
func (s *TokenService) GenerateAccessToken(ctx context.Context, user *entity.User) (string, error) {
	expiredTime := time.Now().Local().Add(10 * time.Minute) //ini untuk mengatur waktu kadaluarsa token
	claims := common.JwtCustomClaims{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Role:  user.Roles,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiredTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)           //ini untuk membuat token
	encodedToken, err := token.SignedString([]byte(s.cfg.JWT.SecretKey)) //ini untuk mengenkripsi token

	if err != nil {
		return "", err
	}
	return encodedToken, nil

}
