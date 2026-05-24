package auth

import (
	"net/http"

	"github.com/vow/app/server/internal/shared/request"
	"github.com/vow/app/server/internal/shared/response"
)

func (h Handler) Login(w http.ResponseWriter, r *http.Request) {
	input, err := request.DecodeAndValidate[LoginRequest](w, r)
	if err != nil {
		response.HandleError(w, err)
		return
	}

	result, err := h.service.Login(r.Context(), input)
	if err != nil {
		response.HandleError(w, err)
		return
	}

	response.OK(w, "logged in successfully", result)
}
