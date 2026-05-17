package middleware

import (
	"net"
	"net/http"
	"sync"
	"time"
)

type rateLimitState struct {
	count int
	reset time.Time
}

func RateLimit(maxRequests int, window time.Duration) Middleware {
	var mu sync.Mutex
	clients := make(map[string]rateLimitState)

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			key := clientIP(r)
			now := time.Now()

			mu.Lock()
			state := clients[key]
			if now.After(state.reset) {
				state = rateLimitState{reset: now.Add(window)}
			}
			state.count++
			clients[key] = state
			allowed := state.count <= maxRequests
			mu.Unlock()

			if !allowed {
				http.Error(w, http.StatusText(http.StatusTooManyRequests), http.StatusTooManyRequests)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

func clientIP(r *http.Request) string {
	if forwardedFor := r.Header.Get("X-Forwarded-For"); forwardedFor != "" {
		return forwardedFor
	}

	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr
	}
	return host
}
