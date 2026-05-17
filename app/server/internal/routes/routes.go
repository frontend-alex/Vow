package routes

import (
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/vow/app/server/internal/auth"
)

func Register(mux *http.ServeMux, db *pgxpool.Pool) {
	auth.RegisterRoutes(mux, auth.NewHandler(auth.NewService(auth.NewRepository(db))))
}
