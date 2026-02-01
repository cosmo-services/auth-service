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

type JwtService struct {
	secret     string
	accessTTL  time.Duration
	refreshTTL time.Duration
}

func NewJwtService(env config.Env) domain.JwtService {
	return &JwtService{
		secret:     env.JwtSecret,
		accessTTL:  env.JwtAccessTTL,
		refreshTTL: env.JwtRefreshTTL,
	}
}

func (s *JwtService) GenerateTokenPair(payload *domain.JwtPayload) (*domain.TokenPair, error) {
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
			Token:   accessToken,
			Expires: accessExpires,
		},
		Refresh: domain.JwtToken{
			Token:   refreshToken,
			Expires: refreshExpires,
		},
	}, nil
}

func (s *JwtService) ValidateToken(tokenStr string) (*domain.JwtPayload, error) {
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

func (s *JwtService) generateToken(payload *domain.JwtPayload, expiresAt time.Time) (string, error) {
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
