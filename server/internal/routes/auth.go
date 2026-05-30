package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
	"vow/server/internal/config"
)

func registerAuthRoutes(router *gin.RouterGroup, cfg *config.Config, log *zap.Logger, db *pgxpool.Pool) {
	router.GET("/health", func(c *gin.Context) {
		log.Debug("auth routes health check", zap.String("app_env", cfg.Server.AppEnv), zap.Bool("db_ready", db != nil))
		c.JSON(http.StatusOK, gin.H{"module": "auth", "status": "ok"})
	})
}
