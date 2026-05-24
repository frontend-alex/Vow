package auth

import (
	"net/http"

	"github.com/vow/app/server/internal/shared/request"
	"github.com/vow/app/server/internal/shared/response"
)

func (h Handler) Register(w http.ResponseWriter, r *http.Request) {
	input, err := request.DecodeAndValidate[RegisterRequest](w, r)
	if err != nil {
		response.HandleError(w, err)
		return
	}

	result, err := h.service.Register(r.Context(), input)
	if err != nil {
		response.HandleError(w, err)
		return
	}

	response.Created(w, "registered successfully", result)
}
