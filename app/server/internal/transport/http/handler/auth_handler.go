package handler

import (
	"net/http"

	"github.com/frontend-alex/Vow/app/server/shared/response"
)

type AuthHandler struct{}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{}
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	response.JSON(w, http.StatusNotImplemented, map[string]string{"message": "register handler not implemented"})
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	response.JSON(w, http.StatusNotImplemented, map[string]string{"message": "login handler not implemented"})
}

func (h *AuthHandler) RefreshToken(w http.ResponseWriter, r *http.Request) {
	response.JSON(w, http.StatusNotImplemented, map[string]string{"message": "refresh token handler not implemented"})
}
