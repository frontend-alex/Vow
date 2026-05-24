package middleware

import (
	"net"
	"net/http"
	"strconv"
	"strings"
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

	go func() {
		ticker := time.NewTicker(window)
		defer ticker.Stop()

		for range ticker.C {
			mu.Lock()
			now := time.Now()

			for key, state := range clients {
				if now.After(state.reset) {
					delete(clients, key)
				}
			}

			mu.Unlock()
		}
	}()

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			key := clientIP(r)
			now := time.Now()

			mu.Lock()

			state := clients[key]
			if now.After(state.reset) {
				state = rateLimitState{
					count: 0,
					reset: now.Add(window),
				}
			}

			state.count++
			clients[key] = state

			allowed := state.count <= maxRequests
			retryAfter := time.Until(state.reset)

			mu.Unlock()

			if !allowed {
				if retryAfter > 0 {
					w.Header().Set("Retry-After", strconv.Itoa(int(retryAfter.Seconds())))
				}

				http.Error(w, http.StatusText(http.StatusTooManyRequests), http.StatusTooManyRequests)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

func clientIP(r *http.Request) string {
	if forwardedFor := r.Header.Get("X-Forwarded-For"); forwardedFor != "" {
		parts := strings.Split(forwardedFor, ",")
		return strings.TrimSpace(parts[0])
	}

	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr
	}

	return host
}
