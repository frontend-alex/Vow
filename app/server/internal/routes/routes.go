package routes

import (
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Router(mux *http.ServeMux, db *pgxpool.Pool) {
	RegisterAuth(mux, db)
}
