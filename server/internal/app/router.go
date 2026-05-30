// Package app wires the HTTP API application.
package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
	"vow/server/internal/config"
	"vow/server/internal/logger"
	"vow/server/internal/routes"
)

// NewRouter creates the API router.
func NewRouter(cfg *config.Config, log *zap.Logger, db *pgxpool.Pool) http.Handler {
	router := gin.New()
	router.Use(logger.Gin(log), logger.GinRecovery(log))

	routes.Register(router, cfg, log, db)

	return router
}
