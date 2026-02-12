package user

type UserRegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type ChangeEmailRequest struct {
	NewEmail string `json:"new_email"`
}
