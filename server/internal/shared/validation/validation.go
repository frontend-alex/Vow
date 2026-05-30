// Package validation provides DTO validation helpers.
package validation

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"

	apperrors "vow/server/internal/shared/errors"
)

var validate = validator.New(validator.WithRequiredStructEnabled())

func init() {
	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		name := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		if name == "" {
			return field.Name
		}
		return name
	})
}

type FieldError struct {
	Field   string `json:"field"`
	Rule    string `json:"rule"`
	Message string `json:"message"`
}

// Struct validates a DTO using `validate` struct tags.
func Struct(dto any) error {
	if err := validate.Struct(dto); err != nil {
		return apperrors.New(
			apperrors.ValidationFailed,
			apperrors.WithCause(err),
			apperrors.WithExtra(map[string]any{"fields": formatErrors(err)}),
		)
	}

	return nil
}

func formatErrors(err error) []FieldError {
	validationErrors, ok := err.(validator.ValidationErrors)
	if !ok {
		return []FieldError{{Message: err.Error()}}
	}

	fields := make([]FieldError, 0, len(validationErrors))
	for _, fieldErr := range validationErrors {
		fields = append(fields, FieldError{
			Field:   fieldErr.Field(),
			Rule:    fieldErr.Tag(),
			Message: fieldMessage(fieldErr),
		})
	}

	return fields
}

func fieldMessage(err validator.FieldError) string {
	switch err.Tag() {
	case "required":
		return err.Field() + " is required."
	case "email":
		return err.Field() + " must be a valid email address."
	case "min":
		return err.Field() + " must be at least " + err.Param() + " characters."
	case "max":
		return err.Field() + " must be at most " + err.Param() + " characters."
	default:
		return err.Field() + " is invalid."
	}
}
