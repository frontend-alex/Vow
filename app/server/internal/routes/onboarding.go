package routes

import (
	"net/http"

	"github.com/vow/app/server/internal/onboarding"
)

func Onboarding(mux *http.ServeMux, deps Dependencies) {
	repository := onboarding.NewRepository(deps.DB)
	service := onboarding.NewService(repository)
	handler := onboarding.NewHandler(service)

	onboarding.Routes(mux, handler)
}
