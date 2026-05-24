package auth

import (
	"net/http"

	"github.com/vow/app/server/internal/shared/apperror"
	"github.com/vow/app/server/internal/shared/response"
)

func (h Handler) Logout(w http.ResponseWriter, r *http.Request) {
	response.HandleError(w, apperror.NotImplemented("AUTH_LOGOUT_NOT_IMPLEMENTED", "logout is not implemented yet"))
}
