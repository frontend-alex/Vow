package auth

import (
	"errors"
	"net/http"

	sharederrors "github.com/vow/app/server/internal/shared/errors"
	"github.com/vow/app/server/internal/shared/request"
	"github.com/vow/app/server/internal/shared/response"
)

func handleRequestError(w http.ResponseWriter, err error) {
	var validationErr request.ValidationError
	if errors.As(err, &validationErr) {
		response.AppErrorWithMessage(w, sharederrors.RequestErrors.ValidationFailed, validationErr.Error())
		return
	}

	if errors.Is(err, request.ErrInvalidJSON) {
		response.AppError(w, sharederrors.RequestErrors.InvalidRequestBody)
		return
	}

	response.AppError(w, sharederrors.RequestErrors.InvalidRequest)
}
