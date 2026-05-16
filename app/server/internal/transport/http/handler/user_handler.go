package handler

import (
	"net/http"

	"github.com/frontend-alex/Vow/app/server/shared/response"
)

type UserHandler struct{}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (h *UserHandler) Me(w http.ResponseWriter, r *http.Request) {
	response.JSON(w, http.StatusNotImplemented, map[string]string{"message": "user handler not implemented"})
}
