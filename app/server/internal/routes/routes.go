package routes

import (
	"log/slog"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/vow/app/server/internal/config"
)

type Dependencies struct {
	DB     *pgxpool.Pool
	Config config.Config
	Logger *slog.Logger
}

func Router(mux *http.ServeMux, deps Dependencies) {
	Authentication(mux, deps)
}
