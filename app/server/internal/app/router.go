package app

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/vow/app/server/internal/config"
	"github.com/vow/app/server/internal/docs"
	"github.com/vow/app/server/internal/middleware"
	"github.com/vow/app/server/internal/routes"
	"github.com/vow/app/server/internal/shared/response"
)

func NewRouter(cfg config.Config, log *slog.Logger, db *pgxpool.Pool) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		response.JSON(w, http.StatusOK, map[string]string{"status": "ok"})
	})
	docs.RegisterRoutes(mux)
	routes.Register(mux, db)

	return middleware.Chain(
		mux,
		middleware.Recover(log),
		middleware.RequestID,
		middleware.Logging(log),
		middleware.SecurityHeaders,
		middleware.CORS(cfg.CORSOrigins),
		middleware.RateLimit(cfg.RateLimitRequests, time.Minute),
	)
}
