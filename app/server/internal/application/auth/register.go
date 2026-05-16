package auth

import "context"

type RegisterInput struct {
	Email       string
	Password    string
	DisplayName string
}

type RegisterOutput struct {
	UserID string
}

type RegisterService struct{}

func NewRegisterService() *RegisterService {
	return &RegisterService{}
}

func (s *RegisterService) Register(ctx context.Context, input RegisterInput) (RegisterOutput, error) {
	return RegisterOutput{}, nil
}
