package auth

import (
	"time"
)

type JwtService interface {
	GenerateTokenPair(payload *JwtPayload) (*TokenPair, error)
	ValidateToken(tokenStr string) (*JwtPayload, error)
}

type TokenPair struct {
	Access  JwtToken `json:"access_token"`
	Refresh JwtToken `json:"refresh_token"`
}

type JwtToken struct {
	Token   string    `json:"token"`
	Expires time.Time `json:"expires"`
}

type JwtPayload struct {
	UserID   string `json:"user_id"`
	IsActive bool   `json:"is_active"`
}
