package tokens

import "time"

type TokenRepository interface {
	Create(token *Token) error
	GetByTokenHash(hash string) (*Token, error)
	DeleteById(tokenId string) error
	DeleteExpired(expireTime time.Time) error
	FindByUserId(userId string, tokenType TokenPurpose) (*Token, error)
}
