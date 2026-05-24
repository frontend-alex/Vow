package auth

import (
	"net/http"

	sharederrors "github.com/vow/app/server/internal/shared/errors"
	"github.com/vow/app/server/internal/shared/response"
)

func (h Handler) Logout(w http.ResponseWriter, r *http.Request) {
	response.AppError(w, sharederrors.GeneralErrors.NotImplemented)
}
