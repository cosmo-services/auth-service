package user

import (
	"main/internal/domain"
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
	eventBus     *domain.EventBus
	emailService EmailService
}

func NewUserService(
	userRepo UserRepository,
	pswdService *password.PasswordService,
	tokenService *tokens.TokenService,
	eventBus *domain.EventBus,
	emailService EmailService,
) *UserService {
	return &UserService{
		userRepo:     userRepo,
		pswdService:  pswdService,
		tokenService: tokenService,
		emailService: emailService,
		eventBus:     eventBus,
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

	if err := s.pswdService.ValidatePassword(password); err != nil {
		return err
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

	s.eventBus.Emit("user.registred", UserRegistredEvent{
		UserID:      user.ID,
		Username:    user.Username,
		Email:       user.Email,
		RegistredAt: user.CreatedAt,
	})

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

	s.eventBus.Emit("user.activated", UserActivateEvent{
		UserID:      user.ID,
		ActivatedAt: time.Now(),
	})

	return nil
}

func (s *UserService) Delete(userId string) error {
	if err := s.tokenService.RevokeAllUserTokens(userId); err != nil {
		return err
	}

	if err := s.userRepo.Delete(userId); err != nil {
		return err
	}

	s.eventBus.Emit("user.deleted", UserDeleteEvent{
		UserID:    userId,
		DeletedAt: time.Now(),
	})

	return nil
}

func (s *UserService) DeleteInactiveUsers() error {
	usersId, err := s.userRepo.DeleteInactiveUsers(time.Now().Add(-ActivateDuration))
	if err != nil {
		return err
	}

	for _, userId := range usersId {
		s.eventBus.Emit("user.deleted", UserDeleteEvent{
			UserID:    userId,
			DeletedAt: time.Now(),
		})
	}

	return nil
}

func (s *UserService) GetUser(userId string) (*User, error) {
	user, err := s.userRepo.GetByID(userId)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) ChangeEmail(userId string, newEmail string) error {
	emailAvailable, err := s.userRepo.IsEmailAvailable(newEmail)
	if err != nil {
		return err
	}
	if !emailAvailable {
		return ErrEmailAlreadyTaken
	}

	user, err := s.userRepo.GetByID(userId)
	if err != nil {
		return err
	}

	if err := user.ChangeEmail(newEmail); err != nil {
		return err
	}

	if err := s.userRepo.Update(user); err != nil {
		return err
	}

	if err := s.sendActivationEmail(user); err != nil {
		return err
	}

	s.eventBus.Emit("user.email.changed", UserChangeEmailEvent{
		UserID:    user.ID,
		NewEmail:  newEmail,
		ChangedAt: time.Now(),
	})

	return nil
}

func (s *UserService) ChangePassword(userId string, newPassword string) error {
	if err := s.pswdService.ValidatePassword(newPassword); err != nil {
		return err
	}

	passwordHash, err := s.pswdService.HashPassword(newPassword)
	if err != nil {
		return err
	}

	user, err := s.userRepo.GetByID(userId)
	if err != nil {
		return err
	}

	if err := user.ChangePassword(passwordHash); err != nil {
		return err
	}

	if err := s.userRepo.Update(user); err != nil {
		return err
	}

	return nil
}

func (s *UserService) ChangeUsername(userId string, newUsername string) error {
	usernameAvailable, err := s.userRepo.IsUsernameAvailable(newUsername)
	if err != nil {
		return err
	}
	if !usernameAvailable {
		return ErrUsernameAlreadyTaken
	}

	user, err := s.userRepo.GetByID(userId)
	if err != nil {
		return err
	}

	if err := user.ChangeUsername(newUsername); err != nil {
		return err
	}

	if err := s.userRepo.Update(user); err != nil {
		return err
	}

	s.eventBus.Emit("user.username.changed", UserChangeUsernameEvent{
		UserID:      user.ID,
		NewUsername: user.Username,
		ChangedAt:   time.Now(),
	})

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
