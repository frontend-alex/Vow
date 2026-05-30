// Package routes registers HTTP routes for the API.
package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
	"vow/server/internal/config"
)

// Register wires all API routes into the Gin engine.
func Register(router *gin.Engine, cfg *config.Config, log *zap.Logger, db *pgxpool.Pool) {
	registerHealthRoutes(router)

	api := router.Group("/v1/api")
	registerAuthRoutes(api.Group("/auth"), cfg, log, db)
	registerUserRoutes(api.Group("/users"), cfg, log, db)
}
