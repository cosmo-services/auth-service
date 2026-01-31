package tokens

type TokenRepository interface {
	Create(token *Token) error
	GetByUserId(userId string) (*Token, error)
	GetByTokenHash(hash string) (*Token, error)
	DeleteById(tokenId string) error
}
