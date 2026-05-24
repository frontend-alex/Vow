package auth

type RegisterRequest struct {
	Email    string `json:"email" validate:"required,email,max=255" sanitize:"trim,lower"`
	Name     string `json:"name" validate:"required,min=2,max=100" sanitize:"trim"`
	Password string `json:"password" validate:"required,min=8,max=72"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email,max=255" sanitize:"trim,lower"`
	Password string `json:"password" validate:"required,min=1,max=72"`
}

type AuthUserResponse struct {
	ID                  int64  `json:"id"`
	Email               string `json:"email"`
	Name                string `json:"name"`
	OnboardingCompleted bool   `json:"onboardingCompleted"`
}

type AuthResponse struct {
	AccessToken string           `json:"accessToken"`
	TokenType   string           `json:"tokenType"`
	User        AuthUserResponse `json:"user"`
}
