package auth

type JwtService interface {
	GenerateTokenPair(userId string) (*TokenPair, error)
	ValidateToken(tokenStr string) (string, error)
}
