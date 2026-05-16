package middleware

import (
	"net/http"
	"slices"

	apperrors "github.com/frontend-alex/Vow/app/server/shared/errors"
	"github.com/frontend-alex/Vow/app/server/shared/response"
)

func RequireRole(allowedRoles ...string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			role := r.Header.Get("X-User-Role")
			if role == "" || !slices.Contains(allowedRoles, role) {
				response.Error(w, r, apperrors.ErrForbidden)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
