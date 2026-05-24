package dto

type RegisterResponse struct {
	User        string `json:"user"`
	AccessToken string `json:"access_token"`
}
