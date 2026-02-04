package auth

import (
	"errors"
	"time"

	"main/internal/config"
	domain "main/internal/domain/auth"

	"github.com/golang-jwt/jwt/v5"
)

type JwtClaims struct {
	Payload *domain.JwtPayload `json:"payload"`
	jwt.RegisteredClaims
}

type JwtClient struct {
	secret     string
	accessTTL  time.Duration
	refreshTTL time.Duration
}

func NewJwtClient(env config.Env) domain.JwtClient {
	return &JwtClient{
		secret:     env.JwtSecret,
		accessTTL:  env.JwtAccessTTL,
		refreshTTL: env.JwtRefreshTTL,
	}
}

func (s *JwtClient) GenerateTokenPair(payload *domain.JwtPayload) (*domain.TokenPair, error) {
	accessExpires := time.Now().Add(s.accessTTL)
	accessToken, err := s.generateToken(payload, accessExpires)
	if err != nil {
		return nil, err
	}

	refreshExpires := time.Now().Add(s.refreshTTL)
	refreshToken, err := s.generateToken(payload, refreshExpires)
	if err != nil {
		return nil, err
	}

	return &domain.TokenPair{
		Access: domain.JwtToken{
			Token:     accessToken,
			ExpiresIn: int64(s.accessTTL.Seconds()),
		},
		Refresh: domain.JwtToken{
			Token:     refreshToken,
			ExpiresIn: int64(s.refreshTTL.Seconds()),
		},
	}, nil
}

func (s *JwtClient) ValidateToken(tokenStr string) (*domain.JwtPayload, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(s.secret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*JwtClaims); ok && token.Valid {
		return claims.Payload, nil
	}

	return nil, errors.New("invalid token")
}

func (s *JwtClient) generateToken(payload *domain.JwtPayload, expiresAt time.Time) (string, error) {
	claims := JwtClaims{
		Payload: payload,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresAt),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.secret))
}
