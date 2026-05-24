package request

import (
	"encoding/json"
	"net/http"

	"github.com/vow/app/server/internal/shared/apperror"
)

const maxRequestBodyBytes = 1 << 20 // 1MB

func DecodeAndValidate[T any](w http.ResponseWriter, r *http.Request) (T, error) {
	var input T

	r.Body = http.MaxBytesReader(w, r.Body, maxRequestBodyBytes)

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&input); err != nil {
		return input, apperror.BadRequest("INVALID_REQUEST_BODY", "invalid request body")
	}

	if err := Sanitize(&input); err != nil {
		return input, apperror.BadRequest("INVALID_REQUEST_BODY", "invalid request body")
	}

	if err := Validate(input); err != nil {
		return input, err
	}

	return input, nil
}
