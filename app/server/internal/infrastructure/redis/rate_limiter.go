package redis

type RateLimiter struct{}

func NewRateLimiter() *RateLimiter {
	return &RateLimiter{}
}
