package auth

import (
	"net/http"

	"github.com/vow/app/server/internal/shared/response"
)

func (h Handler) Logout(w http.ResponseWriter, r *http.Request) {
	response.Error(w, http.StatusNotImplemented, "logout not implemented")
}
