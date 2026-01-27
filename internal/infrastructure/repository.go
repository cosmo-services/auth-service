// internal/repository/postgres/user.go
package infrastructure

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	"main/internal/domain"
	"main/pkg"
)

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
		WHERE username = $1
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

	checkEmailAvailabilityQuery = `
		SELECT 1 FROM users WHERE email = $1 AND is_deleted = false
	`

	checkUsernameAvailabilityQuery = `
		SELECT 1 FROM users WHERE username = $1
	`
)

type userRepository struct {
	db pkg.PostgresDB
}

func NewUserRepository(db pkg.PostgresDB) domain.UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user *domain.User) error {
	if user.ID == "" {
		user.ID = generateID()
	}

	var returnedID string
	err := r.db.QueryRow(
		createUserQuery,
		user.ID,
		user.Username,
		user.PasswordHash,
		strings.ToLower(user.Email),
		user.IsActive,
		user.IsDeleted,
		user.CreatedAt,
		user.UpdatedAt,
	).Scan(&returnedID)

	if err != nil {
		return err
	}

	return nil
}

func (r *userRepository) GetByID(id string) (*domain.User, error) {
	user := &domain.User{}
	err := r.db.QueryRow(getUserByIDQuery, id).Scan(
		&user.ID,
		&user.Username,
		&user.PasswordHash,
		&user.Email,
		&user.IsActive,
		&user.IsDeleted,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrUserNotFound
		}
		return nil, fmt.Errorf("failed to get user by username: %w", err)
	}

	return user, nil
}

func (r *userRepository) GetByEmail(email string) (*domain.User, error) {
	user := &domain.User{}
	err := r.db.QueryRow(getUserByEmailQuery, strings.ToLower(email)).Scan(
		&user.ID,
		&user.Username,
		&user.PasswordHash,
		&user.Email,
		&user.IsActive,
		&user.IsDeleted,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrUserNotFound
		}
		return nil, fmt.Errorf("failed to get user by username: %w", err)
	}

	return user, nil
}

func (r *userRepository) GetByUsername(username string) (*domain.User, error) {
	if username == "" {
		return nil, domain.ErrUserNotFound
	}

	user := &domain.User{}
	err := r.db.QueryRow(getUserByUsernameQuery, username).Scan(
		&user.ID,
		&user.Username,
		&user.PasswordHash,
		&user.Email,
		&user.IsActive,
		&user.IsDeleted,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrUserNotFound
		}
		return nil, fmt.Errorf("failed to get user by username: %w", err)
	}

	return user, nil
}

func (r *userRepository) Update(user *domain.User) error {
	result, err := r.db.Exec(
		updateUserQuery,
		user.ID,
		user.Username,
		user.PasswordHash,
		strings.ToLower(user.Email),
		user.IsActive,
		user.IsDeleted,
		user.UpdatedAt,
	)

	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return domain.ErrUserNotFound
	}

	return nil
}

func (r *userRepository) Delete(userID string) error {
	result, err := r.db.Exec(deleteUserQuery, userID, time.Now().UTC())
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return domain.ErrUserNotFound
	}

	return nil
}

func (r *userRepository) IsEmailAvailable(email string) (bool, error) {
	var exists bool
	err := r.db.QueryRow(checkEmailAvailabilityQuery, strings.ToLower(email)).Scan(&exists)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return true, nil
		}
		return false, fmt.Errorf("failed to check email availability: %w", err)
	}
	return false, nil
}

func (r *userRepository) IsUsernameAvailable(username string) (bool, error) {
	var exists bool
	err := r.db.QueryRow(checkUsernameAvailabilityQuery, username).Scan(&exists)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return true, nil
		}
		return false, fmt.Errorf("failed to check username availability: %w", err)
	}
	return false, nil
}

func generateID() string {
	return fmt.Sprintf("user_%d", time.Now().UnixNano())
}
