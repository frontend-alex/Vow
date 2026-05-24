package auth

import (
	"context"
	"errors"
	"strings"

	sharederrors "github.com/vow/app/server/internal/shared/errors"
	"gorm.io/gorm"
)

type Service struct {
	repository Repository
	jwtSecret  string
}

func NewService(repository Repository, jwtSecret string) Service {
	return Service{
		repository: repository,
		jwtSecret:  jwtSecret,
	}
}

func (s Service) Register(ctx context.Context, input RegisterRequest) (AuthResponse, error) {
	email := strings.TrimSpace(strings.ToLower(input.Email))
	name := strings.TrimSpace(input.Name)

	_, err := s.repository.GetUserByEmail(ctx, email)

	if err == nil {
		return AuthResponse{}, sharederrors.AuthErrors.EmailAlreadyTaken
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return AuthResponse{}, sharederrors.GeneralErrors.InternalServerError
	}

	passwordHash, err := HashPassword(input.Password)
	if err != nil {
		return AuthResponse{}, sharederrors.GeneralErrors.InternalServerError
	}

	user, err := s.repository.CreateUser(ctx, CreateUserParams{
		Email:        email,
		Name:         name,
		PasswordHash: passwordHash,
	})
	if err != nil {
		return AuthResponse{}, sharederrors.GeneralErrors.InternalServerError
	}

	token, err := GenerateAccessToken(user.ID, user.Email, s.jwtSecret)

	if err != nil {
		return AuthResponse{}, sharederrors.GeneralErrors.InternalServerError
	}

	return AuthResponse{
		AccessToken: token,
		TokenType:   "Bearer",
		User: AuthUserResponse{
			ID:                  user.ID,
			Email:               user.Email,
			Name:                user.Name,
			OnboardingCompleted: user.OnboardingCompleted,
		},
	}, nil
}

func (s Service) Login(ctx context.Context, input LoginRequest) (AuthResponse, error) {
	email := strings.TrimSpace(strings.ToLower(input.Email))

	user, err := s.repository.GetUserByEmail(ctx, email)
	if err != nil {
		return AuthResponse{}, sharederrors.AuthErrors.InvalidCredentials
	}

	if err := ComparePassword(user.PasswordHash, input.Password); err != nil {
		return AuthResponse{}, sharederrors.AuthErrors.InvalidCredentials
	}

	token, err := GenerateAccessToken(user.ID, user.Email, s.jwtSecret)
	if err != nil {
		return AuthResponse{}, sharederrors.GeneralErrors.InternalServerError
	}

	return AuthResponse{
		AccessToken: token,
		TokenType:   "Bearer",
		User: AuthUserResponse{
			ID:                  user.ID,
			Email:               user.Email,
			Name:                user.Name,
			OnboardingCompleted: user.OnboardingCompleted,
		},
	}, nil
}
