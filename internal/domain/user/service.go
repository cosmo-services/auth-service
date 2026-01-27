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

func (s *UserService) Register(userRequest *UserCreateRequest) error {
	usernameAvailable, err := s.userRepo.IsUsernameAvailable(userRequest.Username)
	if err != nil {
		return err
	}
	if !usernameAvailable {
		return ErrUsernameAlreadyTaken
	}

	emailAvailable, err := s.userRepo.IsEmailAvailable(userRequest.Email)
	if err != nil {
		return err
	}
	if !emailAvailable {
		return ErrEmailAlreadyTaken
	}

	passwordHash, err := s.pswdService.HashPassword(userRequest.Password)
	if err != nil {
		return err
	}

	now := time.Now()
	user := &User{
		Username:     userRequest.Username,
		PasswordHash: passwordHash,
		Email:        userRequest.Email,
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
