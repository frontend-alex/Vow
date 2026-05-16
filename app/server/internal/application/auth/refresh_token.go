package auth

import "context"

type RefreshTokenInput struct {
	RefreshToken string
}

type RefreshTokenOutput struct {
	AccessToken  string
	RefreshToken string
}

type RefreshTokenService struct{
	
}

func NewRefreshTokenService() *RefreshTokenService {
	return &RefreshTokenService{}
}

func (s *RefreshTokenService) Refresh(ctx context.Context, input RefreshTokenInput) (RefreshTokenOutput, error) {
	return RefreshTokenOutput{}, nil
}
