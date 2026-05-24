package auth

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/vow/app/server/internal/shared/response"
)

func (h Handler) Login(w http.ResponseWriter, r *http.Request) {
	var input LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		response.Error(w, http.StatusBadRequest, "invalid request body")
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
