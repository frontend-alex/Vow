package app

import (
	"context"
	"net/http"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
	"vow/server/internal/config"
	"vow/server/internal/db"
	"vow/server/internal/logger"
)

// App owns the server dependencies and runtime lifecycle.
type App struct {
	config *config.Config
	logger *zap.Logger
	db     *pgxpool.Pool
}

// New initializes application dependencies.
func New() (*App, error) {
	cfg, err := config.Load()
	if err != nil {
		return nil, err
	}

	log, err := logger.New(cfg.Server.AppEnv)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	database, err := db.Connect(ctx, cfg.Database.URL)
	if err != nil {
		_ = log.Sync()
		return nil, err
	}

	log.Info("config loaded", zap.String("app_env", cfg.Server.AppEnv))

	return &App{config: cfg, logger: log, db: database}, nil
}

// Run starts the HTTP server.
func (a *App) Run() error {
	defer func() {
		a.db.Close()
		_ = a.logger.Sync()
	}()

	a.logger.Info("server starting", zap.String("addr", a.config.Server.HTTPAddr))
	err := newServer(a.config, NewRouter(a.config, a.logger, a.db)).ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		return err
	}

	return nil
}
