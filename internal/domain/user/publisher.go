package user

type Publisher interface {
	UserActivated(userId string) error
}
