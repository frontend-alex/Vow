package errors

import (
	stderrors "errors"
	"net/http"
)

func StatusCode(err error) int {
	if stderrors.Is(err, ErrUnauthorized) {
		return http.StatusUnauthorized
	}
	return http.StatusInternalServerError
}
