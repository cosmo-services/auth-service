package tokens

import (
	"time"
)

type TokenService struct {
	tokenRepo      TokenRepository
	tokenHasher    TokenHasher
	tokenGenerator TokenGenerator
}

func NewTokenService(
	tokenRepo TokenRepository,
	tokenHasher TokenHasher,
	tokenGenerator TokenGenerator,
) *TokenService {
	return &TokenService{
		tokenRepo:      tokenRepo,
		tokenHasher:    tokenHasher,
		tokenGenerator: tokenGenerator,
	}
}

func (service *TokenService) UseToken(tokenStr string) (*TokenResult, error) {
	tokenHash, err := service.tokenHasher.HashToken(tokenStr)
	if err != nil {
		return nil, err
	}

	token, err := service.tokenRepo.GetByTokenHash(tokenHash)
	if err != nil {
		return nil, err
	}

	if token.ExpiresAt.Before(time.Now()) {
		return nil, ErrTokenExpired
	}

	if err := service.tokenRepo.DeleteById(token.ID); err != nil {
		return nil, err
	}

	return &TokenResult{
		UserID:    token.UserID,
		TokenType: token.TokenType,
	}, nil
}

func (service *TokenService) RequestToken(userId string, tokenType TokenPurpose) (*TokenRequest, error) {
	existingToken, err := service.tokenRepo.FindByUserId(userId, tokenType)
	if err != nil {
		return nil, err
	}

	if existingToken != nil {
		if err := service.tokenRepo.DeleteById(existingToken.ID); err != nil {
			return nil, err
		}
	}

	tokenStr, err := service.tokenGenerator.GenerateToken()
	if err != nil {
		return nil, err
	}

	tokenHash, err := service.tokenHasher.HashToken(tokenStr)
	if err != nil {
		return nil, err
	}

	token := &Token{
		UserID:    userId,
		Hash:      tokenHash,
		TokenType: tokenType,
		ExpiresAt: time.Now().Add(10 * time.Minute),
		CreatedAt: time.Now(),
	}

	if err := service.tokenRepo.Create(token); err != nil {
		return nil, err
	}

	return &TokenRequest{
		UserID:   userId,
		TokenStr: tokenStr,
	}, nil
}

func (service *TokenService) ClearExpiredTokens() error {
	return service.tokenRepo.DeleteExpired(time.Now())
}

func (service *TokenService) RevokeToken(userId string, tokenType TokenPurpose) error {
	return service.tokenRepo.DeleteByUserId(userId, tokenType)
}

func (service *TokenService) RevokeAllUserTokens(userId string) error {
	return service.tokenRepo.DeleteAllUserTokens(userId)
}
