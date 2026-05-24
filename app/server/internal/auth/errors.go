package auth

import (
	"errors"
	"net/http"

	"github.com/vow/app/server/internal/shared/request"
	"github.com/vow/app/server/internal/shared/response"
)

func handleRequestError(w http.ResponseWriter, err error) {
	var validationErr request.ValidationError
	if errors.As(err, &validationErr) {
		response.Error(w, http.StatusBadRequest, validationErr.Error())
		return
	}

	if errors.Is(err, request.ErrInvalidJSON) {
		response.Error(w, http.StatusBadRequest, "invalid request body")
		return
	}

	response.Error(w, http.StatusBadRequest, "invalid request")
}
