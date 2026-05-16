package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/frontend-alex/Vow/app/server/internal/bootstrap"
)

func main() {
	cfg := bootstrap.LoadConfig()
	logger := bootstrap.NewLogger(cfg.Env)

	db, err := bootstrap.OpenDB(cfg.DatabaseURL)

	if err != nil {
		logger.Error("open database", slog.String("error", err.Error()))
		os.Exit(1)
	}

	if db != nil {
		defer db.Close()
	}

	server := bootstrap.NewServer(cfg, logger, db)

	serverErr := make(chan error, 1)
	go func() {
		logger.Info("server starting", slog.String("addr", cfg.HTTPAddr))
		serverErr <- server.ListenAndServe()
	}()

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	select {

	case err := <-serverErr:
		logger.Error("server stopped", slog.String("error", err.Error()))
		os.Exit(1)
	case sig := <-shutdown:
		logger.Info("shutdown signal received", slog.String("signal", sig.String()))
	}

	ctx, cancel := context.WithTimeout(context.Background(), cfg.ShutdownTimeout)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		logger.Error("graceful shutdown failed", slog.String("error", err.Error()))
		os.Exit(1)
	}
}
