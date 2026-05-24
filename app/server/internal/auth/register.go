package auth

import (
	"errors"
	"net/http"

	"github.com/vow/app/server/internal/shared/request"
	"github.com/vow/app/server/internal/shared/response"
)

func (h Handler) Register(w http.ResponseWriter, r *http.Request) {
	input, err := request.DecodeAndValidate[RegisterRequest](w, r)
	if err != nil {
		handleRequestError(w, err)
		return
	}

	result, err := h.service.Register(r.Context(), input)
	if err != nil {
		if errors.Is(err, ErrEmailAlreadyExists) {
			response.Error(w, http.StatusConflict, "email already exists")
			return
		}

		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	response.OK(w, "registered successfully", result)
}
