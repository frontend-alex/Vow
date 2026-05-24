package middleware

import (
	"log/slog"
	"net/http"

	"github.com/vow/app/server/internal/shared/response"
)

func Recover(log *slog.Logger) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if err := recover(); err != nil {
					log.Error("panic_recovered", "error", err, "path", r.URL.Path)
					response.InternalServerError(w)
				}
			}()
			next.ServeHTTP(w, r)
		})
	}
}
