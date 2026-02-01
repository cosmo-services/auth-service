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

func (s *AuthService) Login(username string, password string) (*TokenPair, error) {
	user, err := s.userRepo.GetByUsername(username)
	if err != nil {
		return nil, err
	}

	if !s.pswdService.CheckPassword(password, user.PasswordHash) {
		return nil, ErrInvalidCredentials
	}

	token, err := s.jwtService.GenerateTokenPair(s.payloadFromUser(user))
	if err != nil {
		return nil, err
	}

	return token, nil
}

func (s *AuthService) Refresh(refreshToken string) (*TokenPair, error) {
	tokenPayload, err := s.jwtService.ValidateToken(refreshToken)
	if err != nil {
		return nil, err
	}

	user, err := s.userRepo.GetByID(tokenPayload.UserID)
	if err != nil {
		return nil, err
	}

	token, err := s.jwtService.GenerateTokenPair(s.payloadFromUser(user))
	if err != nil {
		return nil, err
	}

	return token, nil
}

func (s *AuthService) Validate(accessToken string) (*JwtPayload, error) {
	return s.jwtService.ValidateToken(accessToken)
}

func (s *AuthService) payloadFromUser(user *user_domain.User) *JwtPayload {
	return &JwtPayload{
		UserID:   user.ID,
		IsActive: user.IsActive,
	}
}
