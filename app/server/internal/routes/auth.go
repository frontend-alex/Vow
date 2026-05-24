package routes

import (
	"net/http"

	"github.com/vow/app/server/internal/auth"
)

func Authentication(mux *http.ServeMux, deps Dependencies) {
	repository := auth.NewRepository(deps.DB)

	service := auth.NewService(repository)

	handler := auth.NewHandler(service)

	auth.AuthenticationRoutes(mux, handler)
}
