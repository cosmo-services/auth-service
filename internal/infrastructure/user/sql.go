package infrastructure

const (
	createUserQuery = `
		INSERT INTO users (id, username, password_hash, email, is_active, is_deleted, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id
	`

	getUserByIDQuery = `
		SELECT id, username, password_hash, email, is_active, is_deleted, created_at, updated_at
		FROM users 
		WHERE id = $1
	`

	getUserByEmailQuery = `
		SELECT id, username, password_hash, email, is_active, is_deleted, created_at, updated_at
		FROM users 
		WHERE email = $1 AND is_deleted = false
	`

	getUserByUsernameQuery = `
		SELECT id, username, password_hash, email, is_active, is_deleted, created_at, updated_at
		FROM users 
		WHERE username = $1 AND is_deleted = false
	`

	updateUserQuery = `
		UPDATE users 
		SET username = $2, password_hash = $3, email = $4, 
		    is_active = $5, is_deleted = $6, updated_at = $7
		WHERE id = $1 AND is_deleted = false
	`

	deleteUserQuery = `
		UPDATE users 
		SET is_deleted = true, updated_at = $2 
		WHERE id = $1 AND is_deleted = false
	`

	deleteInactiveUsersQuery = `
		DELETE FROM users 
		WHERE is_active = false 
		AND created_at < $1
		RETURNING id
	`

	checkEmailAvailabilityQuery = `
		SELECT 1 FROM users WHERE email = $1 AND is_deleted = false
	`

	checkUsernameAvailabilityQuery = `
		SELECT 1 FROM users WHERE username = $1
	`
)
