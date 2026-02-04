package tokens

type TokenGenerator interface {
	GenerateToken() (string, error)
}
