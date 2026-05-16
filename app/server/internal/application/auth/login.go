package auth

import "context"

type LoginInput struct {
	Email    string
	Password string
}

type LoginOutput struct {
	AccessToken  string
	RefreshToken string
}

type LoginService struct{}

func NewLoginService() *LoginService {
	return &LoginService{}
}

func (s *LoginService) Login(ctx context.Context, input LoginInput) (LoginOutput, error) {
	return LoginOutput{}, nil
}
