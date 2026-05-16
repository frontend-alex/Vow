package redis

type TokenBlacklist struct{}

func NewTokenBlacklist() *TokenBlacklist {
	return &TokenBlacklist{}
}
