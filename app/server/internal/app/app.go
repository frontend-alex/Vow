package app

import (
	"context"
	"log/slog"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/vow/app/server/internal/config"
	"github.com/vow/app/server/internal/platform/database"
	"github.com/vow/app/server/internal/platform/logger"
)

type App struct {
	config config.Config
	logger *slog.Logger
	db     *pgxpool.Pool
}

func New() (*App, error) {
	cfg := config.Load()
	log := logger.New(cfg.LogLevel)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	db, err := database.NewPostgres(ctx, cfg.DatabaseURL)
	if err != nil {
		return nil, err
	}

	return &App{config: cfg, logger: log, db: db}, nil
}

func (a *App) Run() error {
	defer a.db.Close()

	a.logger.Info("server_starting", "addr", a.config.HTTPAddr)
	return newServer(a.config, NewRouter(a.config, a.logger, a.db)).ListenAndServe()
}
