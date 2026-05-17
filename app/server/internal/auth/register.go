package auth

import (
	"net/http"

	"github.com/vow/app/server/internal/shared/response"
)

func (h Handler) Register(w http.ResponseWriter, r *http.Request) {
	response.JSON(w, http.StatusNotImplemented, map[string]string{"message": "register not implemented"})
}
