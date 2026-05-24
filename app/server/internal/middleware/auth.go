package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/vow/app/server/internal/auth"
	sharederrors "github.com/vow/app/server/internal/shared/errors"
	"github.com/vow/app/server/internal/shared/response"
)

type contextKey string

const UserIDContextKey contextKey = "userID"

func Auth(jwtSecret string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			header := r.Header.Get("Authorization")
			if header == "" {
				response.AppError(w, sharederrors.AuthErrors.Unauthorized)
				return
			}

			tokenString := strings.TrimPrefix(header, "Bearer ")
			if tokenString == header {
				response.AppError(w, sharederrors.AuthErrors.Unauthorized)
				return
			}

			claims, err := auth.ParseAccessToken(tokenString, jwtSecret)
			if err != nil {
				response.AppError(w, sharederrors.JWTErrors.InvalidToken)
				return
			}

			ctx := context.WithValue(r.Context(), UserIDContextKey, claims.UserID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func UserIDFromContext(ctx context.Context) (int64, bool) {
	userID, ok := ctx.Value(UserIDContextKey).(int64)
	return userID, ok
}
