package auth

type JwtService interface {
	GenerateToken(userId string) (*AccessToken, error)
	ValidateToken(tokenStr string) (string, error)
}
