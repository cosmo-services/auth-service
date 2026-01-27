package infrastructure

import (
	"errors"
	"strings"

	"main/internal/domain/password"

	"golang.org/x/crypto/bcrypt"
)

const (
	DefaultCost = 12
)

type BcryptPasswordHasher struct {
	cost int
}

func NewBcryptPasswordHasher(cost int) password.PasswordHasher {
	return &BcryptPasswordHasher{
		cost: DefaultCost,
	}
}

func (h *BcryptPasswordHasher) HashPassword(password string) (string, error) {
	if password == "" {
		return "", errors.New("password cannot be empty")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), h.cost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func (h *BcryptPasswordHasher) CheckPassword(password string, hash string) bool {
	if password == "" || hash == "" {
		return false
	}

	hash = strings.TrimSpace(hash)

	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err == nil
}
