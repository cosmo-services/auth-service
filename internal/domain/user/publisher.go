package user

type Publisher interface {
	UserActivated(userId string) error
	UserDeleted(userId string) error
}
