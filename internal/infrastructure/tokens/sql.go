package tokens

const (
	createTokenQuery = `
		INSERT INTO tokens (id, user_id, hash, token_type, expires_at, created_at)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id
	`

	getTokenByHashQuery = `
		SELECT id, user_id, hash, token_type, expires_at, created_at
		FROM tokens 
		WHERE hash = $1
	`

	deleteTokenByIDQuery = `
		DELETE FROM tokens 
		WHERE id = $1
	`

	deleteExpiredTokensQuery = `
		DELETE FROM tokens 
		WHERE expires_at < $1
	`

	deleteTokenByUserIDAndTypeQuery = `
		DELETE FROM tokens 
		WHERE user_id = $1 AND token_type = $2
	`

	findTokenByUserIDAndTypeQuery = `
		SELECT id, user_id, hash, token_type, expires_at, created_at
		FROM tokens 
		WHERE user_id = $1 AND token_type = $2
		ORDER BY created_at DESC
		LIMIT 1
	`
)
