package auth

type RegisterRequest struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
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
