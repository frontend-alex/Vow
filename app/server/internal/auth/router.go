package auth

import "net/http"

func RegisterRoutes(mux *http.ServeMux, handler Handler) {
	mux.HandleFunc("POST /v1/auth/login", handler.Login)
	mux.HandleFunc("POST /v1/auth/register", handler.Register)
	mux.HandleFunc("POST /v1/auth/logout", handler.Logout)
}
