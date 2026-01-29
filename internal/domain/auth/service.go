package auth

import (
	"main/internal/domain/password"
	user_domain "main/internal/domain/user"
)

type AuthService struct {
	jwtService  JwtService
	pswdService *password.PasswordService
	userRepo    user_domain.UserRepository
}

func NewAuthService(
	jwtService JwtService,
	pswdService *password.PasswordService,
	userRepo user_domain.UserRepository,
) *AuthService {
	return &AuthService{
		jwtService:  jwtService,
		pswdService: pswdService,
		userRepo:    userRepo,
	}
}

func (s *AuthService) Login(username string, password string) (*AccessToken, error) {
	user, err := s.userRepo.GetByUsername(username)
	if err != nil {
		return nil, err
	}

	if !user.IsActive {
		return nil, ErrInactiveUser
	}

	if !s.pswdService.CheckPassword(password, user.PasswordHash) {
		return nil, ErrInvalidCredentials
	}

	token, err := s.jwtService.GenerateToken(user.ID)
	if err != nil {
		return nil, err
	}

	return token, nil
}
