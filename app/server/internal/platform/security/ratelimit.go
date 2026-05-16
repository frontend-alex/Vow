package security

import (
	"net"
	"net/http"
	"sync"
	"time"

	apperrors "github.com/frontend-alex/Vow/app/server/shared/errors"
	"github.com/frontend-alex/Vow/app/server/shared/response"
)

type visitor struct {
	count     int
	resetTime time.Time
}

func RateLimit(requestsPerMinute int) func(http.Handler) http.Handler {
	if requestsPerMinute <= 0 {
		requestsPerMinute = 120
	}

	var mu sync.Mutex
	visitors := map[string]*visitor{}

	go func() {
		for range time.Tick(time.Minute) {
			mu.Lock()
			now := time.Now()
			for ip, v := range visitors {
				if now.After(v.resetTime.Add(time.Minute)) {
					delete(visitors, ip)
				}
			}
			mu.Unlock()
		}
	}()

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ip := clientIP(r)
			now := time.Now()

			mu.Lock()
			v := visitors[ip]
			if v == nil || now.After(v.resetTime) {
				v = &visitor{resetTime: now.Add(time.Minute)}
				visitors[ip] = v
			}
			v.count++
			limited := v.count > requestsPerMinute
			mu.Unlock()

			if limited {
				response.Error(w, r, apperrors.ErrRateLimited)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

func clientIP(r *http.Request) string {
	if forwarded := r.Header.Get("X-Forwarded-For"); forwarded != "" {
		return forwarded
	}

	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr
	}
	return host
}
