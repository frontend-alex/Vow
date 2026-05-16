package http

import (
	"database/sql"
	"log/slog"
	"net/http"

	"github.com/frontend-alex/Vow/app/server/internal/platform/security"
	"github.com/frontend-alex/Vow/app/server/internal/transport/http/docs"
	"github.com/frontend-alex/Vow/app/server/internal/transport/http/handler"
	"github.com/frontend-alex/Vow/app/server/internal/transport/http/middleware"
	"github.com/frontend-alex/Vow/app/server/shared/response"
)

type RouterConfig struct {
	AllowedOrigins []string
	RateLimitRPM   int
}

func NewRouter(cfg RouterConfig, logger *slog.Logger, db *sql.DB) http.Handler {
	mux := http.NewServeMux()
	authHandler := handler.NewAuthHandler()
	userHandler := handler.NewUserHandler()

	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		payload := map[string]string{"status": "ok"}
		if db == nil {
			payload["database"] = "not_configured"
		} else if err := db.PingContext(r.Context()); err != nil {
			response.Error(w, r, err)
			return
		} else {
			payload["database"] = "ok"
		}

		response.JSON(w, http.StatusOK, payload)
	})
	mux.HandleFunc("POST /auth/register", authHandler.Register)
	mux.HandleFunc("POST /auth/login", authHandler.Login)
	mux.HandleFunc("POST /auth/refresh", authHandler.RefreshToken)

	mux.Handle("GET /users/me", middleware.RequireAuth(http.HandlerFunc(userHandler.Me)))

	mux.HandleFunc("GET /docs", swaggerUI)
	mux.HandleFunc("GET /docs/", swaggerUI)
	mux.HandleFunc("GET /docs/openapi.yaml", openAPIYAML)

	var handler http.Handler = mux
	
	handler = security.RateLimit(cfg.RateLimitRPM)(handler)
	handler = security.CORS(cfg.AllowedOrigins)(handler)
	handler = security.Headers(handler)
	handler = middleware.RequestID(handler)
	handler = middleware.Logging(logger)(handler)
	handler = middleware.Recover(logger)(handler)

	return handler
}

func swaggerUI(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(docs.SwaggerHTML))
}

func openAPIYAML(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/yaml; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(docs.OpenAPIYAML)
}
