package auth

import "time"

type TokenPair struct {
	AccessToken           string
	RefreshToken          string
	AccessTokenExpiresAt  time.Time
	RefreshTokenExpiresAt time.Time
}

type Claims struct {
	Subject string
	Role    string
	Expiry  time.Time
}
