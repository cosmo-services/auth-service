package tokens

import "time"

type TokenPurpose string

const (
	PurposeVerifyEmail   TokenPurpose = "verify_email"
	PurposePasswordReset TokenPurpose = "password_reset"
)

type Token struct {
	ID        string       `json:"id"`
	UserID    string       `json:"user_id"`
	Hash      string       `json:"hash"`
	TokenType TokenPurpose `json:"token_type"`
	ExpiresAt time.Time    `json:"expires_at"`
	CreatedAt time.Time    `json:"created_at"`
}

type TokenRequest struct {
	UserID   string `json:"user_id"`
	TokenStr string `json:"token"`
}

type TokenResult struct {
	UserID    string       `json:"user_id"`
	TokenType TokenPurpose `json:"token_type"`
}
