package domain

type PasswordHasher interface {
	HashPassword(password string) (string, error)
	CheckPassword(password string, hash string) bool
}
