package user

import (
	"main/internal/domain/password"
	"time"
)

type UserService struct {
	userRepo    UserRepository
	pswdService *password.PasswordService
}

func NewUserService(
	userRepo UserRepository,
	pswdService *password.PasswordService,
) *UserService {
	return &UserService{
		userRepo:    userRepo,
		pswdService: pswdService,
	}
}

func (s *UserService) Register(username string, password string, email string) error {
	usernameAvailable, err := s.userRepo.IsUsernameAvailable(username)
	if err != nil {
		return err
	}
	if !usernameAvailable {
		return ErrUsernameAlreadyTaken
	}

	emailAvailable, err := s.userRepo.IsEmailAvailable(email)
	if err != nil {
		return err
	}
	if !emailAvailable {
		return ErrEmailAlreadyTaken
	}

	passwordHash, err := s.pswdService.HashPassword(password)
	if err != nil {
		return err
	}

	now := time.Now()
	user := &User{
		Username:     username,
		PasswordHash: passwordHash,
		Email:        email,
		IsActive:     false,
		IsDeleted:    false,
		CreatedAt:    now,
		UpdatedAt:    now,
	}

	if err := s.userRepo.Create(user); err != nil {
		return err
	}

	return nil
}
