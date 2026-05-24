package request

import (
	"encoding/json"
	"errors"
	"net/http"
)

const maxRequestBodyBytes = 1 << 20 // 1MB

var ErrInvalidJSON = errors.New("invalid request body")

func DecodeAndValidate[T any](w http.ResponseWriter, r *http.Request) (T, error) {
	var input T

	r.Body = http.MaxBytesReader(w, r.Body, maxRequestBodyBytes)

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&input); err != nil {
		return input, ErrInvalidJSON
	}

	if err := Sanitize(&input); err != nil {
		return input, err
	}

	if err := Validate(input); err != nil {
		return input, err
	}

	return input, nil
}
