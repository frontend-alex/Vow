package auth

import (
	"net/http"

	"github.com/vow/app/server/internal/shared/response"
)

func (h Handler) Login(w http.ResponseWriter, r *http.Request) {
	response.JSON(w, http.StatusNotImplemented, map[string]string{"message": "login not implemented"})
}
