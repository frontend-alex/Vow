package auth

import (
	"context"
	"errors"
	"strings"

	"github.com/vow/app/server/internal/shared/apperror"
	"gorm.io/gorm"
)

var (
	ErrInvalidCredentials = apperror.Unauthorized("AUTH_INVALID_CREDENTIALS", "invalid email or password")
	ErrEmailAlreadyExists = apperror.Conflict("AUTH_EMAIL_ALREADY_EXISTS", "email already exists")
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
		return AuthResponse{}, ErrEmailAlreadyExists
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return AuthResponse{}, err
	}

	passwordHash, err := HashPassword(input.Password)
	if err != nil {
		return AuthResponse{}, apperror.Internal()
	}

	user, err := s.repository.CreateUser(ctx, CreateUserParams{
		Email:        email,
		Name:         name,
		PasswordHash: passwordHash,
	})
	if err != nil {
		return AuthResponse{}, err
	}

	token, err := GenerateAccessToken(user.ID, user.Email, s.jwtSecret)
	if err != nil {
		return AuthResponse{}, apperror.Internal()
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
		return AuthResponse{}, ErrInvalidCredentials
	}

	if err := ComparePassword(user.PasswordHash, input.Password); err != nil {
		return AuthResponse{}, ErrInvalidCredentials
	}

	token, err := GenerateAccessToken(user.ID, user.Email, s.jwtSecret)
	if err != nil {
		return AuthResponse{}, apperror.Internal()
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
