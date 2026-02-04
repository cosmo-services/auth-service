package tokens

import (
	"crypto/sha256"
	"encoding/hex"
	domain "main/internal/domain/tokens"
)

type Sha256Hasher struct{}

func NewSha256Hasher() domain.TokenHasher {
	return &Sha256Hasher{}
}

func (h *Sha256Hasher) HashToken(tokenStr string) (string, error) {
	hash := sha256.Sum256([]byte(tokenStr))
	return hex.EncodeToString(hash[:]), nil
}
