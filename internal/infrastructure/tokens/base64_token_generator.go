package tokens

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	domain "main/internal/domain/tokens"
)

type Base64TokenGenerator struct{}

func NewBase64TokenGenerator() domain.TokenGenerator {
	return &Base64TokenGenerator{}
}

func (g *Base64TokenGenerator) GenerateToken() (string, error) {
	bytes := make([]byte, 32)
	if _, err := rand.Read(bytes); err != nil {
		return "", fmt.Errorf("failed to generate random token: %w", err)
	}
	return base64.RawURLEncoding.EncodeToString(bytes), nil
}
