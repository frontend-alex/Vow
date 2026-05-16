package bootstrap

import (
	"database/sql"
	"log/slog"
	"net/http"

	transporthttp "github.com/frontend-alex/Vow/app/server/internal/transport/http"
)

func NewServer(cfg Config, logger *slog.Logger, db *sql.DB) *http.Server {
	return &http.Server{
		Addr:         cfg.HTTPAddr,
		Handler:      transporthttp.NewRouter(transporthttp.RouterConfig{ AllowedOrigins: cfg.AllowedOrigins, RateLimitRPM: cfg.RateLimitRPM }, logger, db),
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
		IdleTimeout:  cfg.IdleTimeout,
	}
}
