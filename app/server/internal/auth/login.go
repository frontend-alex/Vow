package auth

import (
	"errors"
	"net/http"

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
		if errors.Is(err, ErrInvalidCredentials) {
			response.Error(w, http.StatusUnauthorized, "invalid email or password")
			return
		}

		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	response.OK(w, "logged in successfully", result)
}
