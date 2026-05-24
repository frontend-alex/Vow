package auth

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/vow/app/server/internal/shared/response"
)

func (h Handler) Register(w http.ResponseWriter, r *http.Request) {
	var input RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		response.Error(w, http.StatusBadRequest, "invalid request body")
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
