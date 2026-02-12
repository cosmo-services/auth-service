package auth

import (
	"main/internal/domain/password"
	user_domain "main/internal/domain/user"
)

type AuthService struct {
	jwtClient   JwtClient
	pswdService *password.PasswordService
	userRepo    user_domain.UserRepository
}

func NewAuthService(
	jwtClient JwtClient,
	pswdService *password.PasswordService,
	userRepo user_domain.UserRepository,
) *AuthService {
	return &AuthService{
		jwtClient:   jwtClient,
		pswdService: pswdService,
		userRepo:    userRepo,
	}
}

func (s *AuthService) Login(username string, password string) (*TokenPair, error) {
	user, err := s.userRepo.GetByUsername(username)
	if err != nil {
		return nil, ErrInvalidCredentials
	}

	if !s.pswdService.CheckPassword(password, user.PasswordHash) {
		return nil, ErrInvalidCredentials
	}

	token, err := s.jwtClient.GenerateTokenPair(s.payloadFromUser(user))
	if err != nil {
		return nil, err
	}

	return token, nil
}

func (s *AuthService) Refresh(refreshToken string) (*TokenPair, error) {
	tokenPayload, err := s.jwtClient.ValidateToken(refreshToken)
	if err != nil {
		return nil, err
	}

	user, err := s.userRepo.GetByID(tokenPayload.UserID)
	if err != nil {
		return nil, err
	}

	token, err := s.jwtClient.GenerateTokenPair(s.payloadFromUser(user))
	if err != nil {
		return nil, err
	}

	return token, nil
}

func (s *AuthService) Validate(accessToken string) (*JwtPayload, error) {
	return s.jwtClient.ValidateToken(accessToken)
}

func (s *AuthService) payloadFromUser(user *user_domain.User) *JwtPayload {
	return &JwtPayload{
		UserID:   user.ID,
		IsActive: user.IsActive,
	}
}
