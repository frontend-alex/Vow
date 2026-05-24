package routes

import (
	"log/slog"
	"net/http"

	"github.com/vow/app/server/internal/config"
	"gorm.io/gorm"
)

type Dependencies struct {
	DB     *gorm.DB
	Config config.Config
	Logger *slog.Logger
}

func Router(mux *http.ServeMux, deps Dependencies) {
	Authentication(mux, deps)
}
