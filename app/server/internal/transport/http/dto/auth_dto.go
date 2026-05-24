package dto

type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type RegisterResponse struct {
	User        string `json:"user"`
	AccessToken string `json:"access_token"`
}
