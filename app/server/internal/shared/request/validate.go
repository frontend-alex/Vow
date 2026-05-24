package request

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/vow/app/server/internal/shared/apperror"
)

var validate = validator.New()

func init() {
	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		name := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
}

func Validate(input any) error {
	err := validate.Struct(input)
	if err == nil {
		return nil
	}

	var validationErrors validator.ValidationErrors
	if !errors.As(err, &validationErrors) {
		return apperror.BadRequest("INVALID_REQUEST", "invalid request")
	}

	fields := make([]apperror.FieldError, 0, len(validationErrors))
	for _, fieldError := range validationErrors {
		fields = append(fields, apperror.FieldError{
			Field:   fieldError.Field(),
			Message: validationMessage(fieldError),
		})
	}

	return apperror.Validation(fields)
}

func validationMessage(err validator.FieldError) string {
	switch err.Tag() {
	case "required":
		return "is required"
	case "email":
		return "must be a valid email address"
	case "min":
		return fmt.Sprintf("must be at least %s characters", err.Param())
	case "max":
		return fmt.Sprintf("must be at most %s characters", err.Param())
	case "oneof":
		return fmt.Sprintf("must be one of: %s", err.Param())
	default:
		return "is invalid"
	}
}
