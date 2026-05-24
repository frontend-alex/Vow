package auth

import (
	"net/http"

	sharederrors "github.com/vow/app/server/internal/shared/errors"
	"github.com/vow/app/server/internal/shared/request"
	"github.com/vow/app/server/internal/shared/response"
)

func (h Handler) Login(w http.ResponseWriter, r *http.Request) {
	input, err := request.DecodeAndValidate[LoginRequest](w, r)
	if err != nil {
		handleRequestError(w, err)
		return
	}

	result, err := h.service.Login(r.Context(), input)
	if err != nil {
		if apiError, ok := sharederrors.FromError(err); ok {
			response.AppError(w, apiError)
			return
		}

		response.InternalServerError(w)
		return
	}

	response.OK(w, "logged in successfully", result)
}
