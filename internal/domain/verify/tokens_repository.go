package verify

type VerifyTokenRepository interface {
	Create(token *VerifyToken) error
	GetByUserId(userId string) (*VerifyToken, error)
	DeleteByUserID(userId string) error
}
