package auth

import "time"

type Session struct {
	ID           string
	UserID       string
	RefreshToken string
	ExpiresAt    time.Time
	CreatedAt    time.Time
	RevokedAt    *time.Time
}

func (s Session) Active(now time.Time) bool {
	return s.RevokedAt == nil && now.Before(s.ExpiresAt)
}
