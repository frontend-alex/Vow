package middleware

import (
	"log/slog"
	"net/http"

	apperrors "github.com/frontend-alex/Vow/app/server/shared/errors"
	"github.com/frontend-alex/Vow/app/server/shared/response"
)

func Recover(logger *slog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if recovered := recover(); recovered != nil {
					logger.Error("panic recovered", slog.Any("panic", recovered), slog.String("path", r.URL.Path))
					response.Error(w, r, apperrors.ErrInternal)
				}
			}()

			next.ServeHTTP(w, r)
		})
	}
}
