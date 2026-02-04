package user

type EmailService interface {
	SendToken(token string, email string) error
}
