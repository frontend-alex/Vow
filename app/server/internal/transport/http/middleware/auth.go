package middleware

import (
	"net/http"
	"strings"

	apperrors "github.com/frontend-alex/Vow/app/server/shared/errors"
	"github.com/frontend-alex/Vow/app/server/shared/response"
)

func RequireAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")
		
		if !strings.HasPrefix(header, "Bearer ") || strings.TrimSpace(strings.TrimPrefix(header, "Bearer ")) == "" {
			response.Error(w, r, apperrors.ErrUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
