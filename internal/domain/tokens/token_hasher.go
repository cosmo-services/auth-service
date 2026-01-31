package tokens

type TokenHasher interface {
	HashToken(tokenStr string) (string, error)
}
