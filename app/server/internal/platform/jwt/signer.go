package jwt

import "time"

type Signer struct {
	secret []byte
}

func NewSigner(secret string) *Signer {
	return &Signer{secret: []byte(secret)}
}

func (s *Signer) Sign(subject string, ttl time.Duration) (string, time.Time, error) {
	expiresAt := time.Now().Add(ttl)
	return "", expiresAt, nil
}
