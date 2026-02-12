package auth

type JwtClient interface {
	GenerateTokenPair(payload *JwtPayload) (*TokenPair, error)
	ValidateToken(tokenStr string) (*JwtPayload, error)
}

type TokenPair struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int64  `json:"expires_in"`
}

type JwtPayload struct {
	UserID   string `json:"user_id"`
	IsActive bool   `json:"is_active"`
}
