// Package logger provides centralized application logging.
package logger

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const productionEnv = "production"

func New(appEnv string) (*zap.Logger, error) {
	if isProduction(appEnv) {
		return zap.NewProduction(zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
	}

	encoderConfig := zap.NewDevelopmentEncoderConfig()
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339)
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder

	config := zap.Config{
		Level:             zap.NewAtomicLevelAt(zap.DebugLevel),
		Development:       true,
		DisableCaller:     false,
		DisableStacktrace: false,
		Sampling:          nil,
		Encoding:          "console",
		EncoderConfig:     encoderConfig,
		OutputPaths:       []string{"stdout"},
		ErrorOutputPaths:  []string{"stderr"},
	}

	return config.Build(zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
}

func isProduction(appEnv string) bool {
	return strings.EqualFold(strings.TrimSpace(appEnv), productionEnv)
}

// Gin returns request logging middleware backed by Zap.
func Gin(log *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		c.Next()

		if raw != "" {
			path += "?" + raw
		}

		status := c.Writer.Status()
		fields := []zap.Field{
			zap.Int("status", status),
			zap.String("method", c.Request.Method),
			zap.String("path", path),
			zap.String("ip", c.ClientIP()),
			zap.String("user_agent", c.Request.UserAgent()),
			zap.Duration("latency", time.Since(start)),
			zap.Int("bytes", c.Writer.Size()),
		}

		if len(c.Errors) > 0 {
			fields = append(fields, zap.String("errors", c.Errors.String()))
		}

		switch {
		case status >= http.StatusInternalServerError:
			log.Error("http request completed", fields...)
		case status >= http.StatusBadRequest:
			log.Warn("http request completed", fields...)
		default:
			log.Info("http request completed", fields...)
		}
	}
}

// GinRecovery returns panic recovery middleware backed by Zap.
func GinRecovery(log *zap.Logger) gin.HandlerFunc {
	return gin.CustomRecovery(func(c *gin.Context, recovered any) {
		log.Error(
			"panic recovered",
			zap.Any("panic", recovered),
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.URL.Path),
			zap.String("ip", c.ClientIP()),
		)
		c.AbortWithStatus(http.StatusInternalServerError)
	})
}
