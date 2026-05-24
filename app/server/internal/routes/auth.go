package routes

import (
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/vow/app/server/internal/auth"
)

func RegisterAuth(mux *http.ServeMux, db *pgxpool.Pool) {
	handler := auth.NewHandler(auth.NewService(auth.NewRepository(db)))

	auth.RegisterRoutes(mux, handler)
}
