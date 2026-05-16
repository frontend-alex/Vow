package user

import "context"

type GetUserInput struct {
	UserID string
}

type GetUserOutput struct {
	UserID      string
	Email       string
	DisplayName string
}

type GetUserService struct{}

func NewGetUserService() *GetUserService {
	return &GetUserService{}
}

func (s *GetUserService) Get(ctx context.Context, input GetUserInput) (GetUserOutput, error) {
	return GetUserOutput{}, nil
}
