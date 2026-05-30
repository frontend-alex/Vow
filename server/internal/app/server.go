package app

import (
	"net/http"
	"time"

	"vow/server/internal/config"
)

func newServer(cfg *config.Config, handler http.Handler) *http.Server {
	return &http.Server{
		Addr:              cfg.Server.HTTPAddr,
		Handler:           handler,
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       time.Minute,
	}
}
