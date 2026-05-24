package auth

import "net/http"

func AuthenticationRoutes(mux *http.ServeMux, handler Handler) {
	mux.HandleFunc("POST /v1/api/auth/login", handler.Login)
	mux.HandleFunc("POST /v1/api/auth/register", handler.Register)
	mux.HandleFunc("POST /v1/api/auth/logout", handler.Logout)
}
