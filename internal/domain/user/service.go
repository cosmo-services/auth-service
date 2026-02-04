package user

import (
	"main/internal/domain/password"
	"main/internal/domain/tokens"
	"time"
)

const (
	ActivateDuration = 7 * 24 * time.Hour
)

type UserService struct {
	userRepo     UserRepository
	pswdService  *password.PasswordService
	tokenService *tokens.TokenService
	publisher    Publisher
	emailService EmailService
}

func NewUserService(
	userRepo UserRepository,
	pswdService *password.PasswordService,
	tokenService *tokens.TokenService,
	publisher Publisher,
	emailService EmailService,
) *UserService {
	return &UserService{
		userRepo:     userRepo,
		pswdService:  pswdService,
		tokenService: tokenService,
		publisher:    publisher,
		emailService: emailService,
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

	if err := s.sendActivationEmail(user); err != nil {
		return err
	}

	return nil
}

func (s *UserService) ResendActivation(userId string) error {
	user, err := s.userRepo.GetByID(userId)
	if err != nil {
		return err
	}

	if err := user.ValidateActivation(); err != nil {
		return err
	}

	if err := s.sendActivationEmail(user); err != nil {
		return err
	}

	return nil
}

func (s *UserService) Activate(tokenStr string) error {
	token, err := s.tokenService.UseToken(tokenStr)
	if err != nil {
		return err
	}

	user, err := s.userRepo.GetByID(token.UserID)
	if err != nil {
		return err
	}

	if err := user.Activate(); err != nil {
		return err
	}

	if err := s.userRepo.Update(user); err != nil {
		return err
	}

	if err := s.publisher.UserActivated(token.UserID); err != nil {
		return err
	}

	return nil
}

func (s *UserService) Delete(userId string) error {
	user, err := s.userRepo.GetByID(userId)
	if err != nil {
		return err
	}

	if err := s.tokenService.RevokeAllUserTokens(userId); err != nil {
		return err
	}

	if err := s.userRepo.Delete(userId); err != nil {
		return err
	}

	if user.IsActive {
		if err := s.publisher.UserDeleted(userId); err != nil {
			return err
		}
	}

	return nil
}

func (s *UserService) DeleteInactiveUsers() error {
	if err := s.userRepo.DeleteInactiveUsers(time.Now().Add(-ActivateDuration)); err != nil {
		return err
	}

	return nil
}

func (s *UserService) sendActivationEmail(user *User) error {
	token, err := s.tokenService.RequestToken(user.ID, tokens.PurposeVerifyEmail)
	if err != nil {
		return err
	}

	if err := s.emailService.SendToken(token.TokenStr, user.Email); err != nil {
		return err
	}

	return nil
}
