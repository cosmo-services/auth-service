package tokens

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	domain "main/internal/domain/tokens"
	"main/pkg"
)

type tokenRepository struct {
	db pkg.PostgresDB
}

func NewTokenRepository(db pkg.PostgresDB) domain.TokenRepository {
	return &tokenRepository{db: db}
}

func (r *tokenRepository) Create(token *domain.Token) error {
	if token.ID == "" {
		token.ID = generateTokenID()
	}

	if token.CreatedAt.IsZero() {
		token.CreatedAt = time.Now().UTC()
	}

	var returnedID string
	err := r.db.QueryRow(
		createTokenQuery,
		token.ID,
		token.UserID,
		token.Hash,
		token.TokenType,
		token.ExpiresAt,
		token.CreatedAt,
	).Scan(&returnedID)

	if err != nil {
		return fmt.Errorf("failed to create token: %w", err)
	}

	return nil
}

func (r *tokenRepository) GetByTokenHash(hash string) (*domain.Token, error) {
	token := &domain.Token{}
	err := r.db.QueryRow(getTokenByHashQuery, hash).Scan(
		&token.ID,
		&token.UserID,
		&token.Hash,
		&token.TokenType,
		&token.ExpiresAt,
		&token.CreatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrTokenNotFound
		}
		return nil, fmt.Errorf("failed to get token by hash: %w", err)
	}

	return token, nil
}

func (r *tokenRepository) DeleteById(tokenId string) error {
	result, err := r.db.Exec(deleteTokenByIDQuery, tokenId)
	if err != nil {
		return fmt.Errorf("failed to delete token: %w", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return domain.ErrTokenNotFound
	}

	return nil
}

func (r *tokenRepository) DeleteExpired(expireTime time.Time) error {
	_, err := r.db.Exec(deleteExpiredTokensQuery, expireTime)
	if err != nil {
		return fmt.Errorf("failed to delete expired tokens: %w", err)
	}

	return nil
}

func (r *tokenRepository) DeleteByUserId(userId string, tokenType domain.TokenPurpose) error {
	_, err := r.db.Exec(deleteTokenByUserIDAndTypeQuery, userId, tokenType)
	if err != nil {
		return fmt.Errorf("failed to delete user's token: %w", err)
	}

	return nil
}

func (r *tokenRepository) DeleteAllUserTokens(userId string) error {
	_, err := r.db.Exec(deleteTokensByUserIDQuery, userId)
	if err != nil {
		return fmt.Errorf("failed to delete user's token: %w", err)
	}
	return nil
}

func (r *tokenRepository) FindByUserId(userId string, tokenType domain.TokenPurpose) (*domain.Token, error) {
	token := &domain.Token{}
	err := r.db.QueryRow(findTokenByUserIDAndTypeQuery, userId, tokenType).Scan(
		&token.ID,
		&token.UserID,
		&token.Hash,
		&token.TokenType,
		&token.ExpiresAt,
		&token.CreatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to find token by user id and type: %w", err)
	}

	return token, nil
}

func generateTokenID() string {
	return fmt.Sprintf("token_%d", time.Now().UnixNano())
}
