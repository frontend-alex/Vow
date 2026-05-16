package user

import "context"

type UpdateUserInput struct {
	UserID      string
	DisplayName string
}

type UpdateUserOutput struct {
	UserID      string
	DisplayName string
}

type UpdateUserService struct{}

func NewUpdateUserService() *UpdateUserService {
	return &UpdateUserService{}
}

func (s *UpdateUserService) Update(ctx context.Context, input UpdateUserInput) (UpdateUserOutput, error) {
	return UpdateUserOutput{}, nil
}
