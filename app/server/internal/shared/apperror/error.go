package apperror

import "net/http"

type AppError struct {
	Code    string
	Status  int
	Message string
	Fields  []FieldError
}

type FieldError struct {
	Field   string
	Message string
}

func (e AppError) Error() string {
	return e.Message
}

func BadRequest(code string, message string) AppError {
	return AppError{Code: code, Status: http.StatusBadRequest, Message: message}
}

func Unauthorized(code string, message string) AppError {
	return AppError{Code: code, Status: http.StatusUnauthorized, Message: message}
}

func Forbidden(code string, message string) AppError {
	return AppError{Code: code, Status: http.StatusForbidden, Message: message}
}

func NotFound(code string, message string) AppError {
	return AppError{Code: code, Status: http.StatusNotFound, Message: message}
}

func Conflict(code string, message string) AppError {
	return AppError{Code: code, Status: http.StatusConflict, Message: message}
}

func NotImplemented(code string, message string) AppError {
	return AppError{Code: code, Status: http.StatusNotImplemented, Message: message}
}

func Internal() AppError {
	return AppError{Code: "INTERNAL_SERVER_ERROR", Status: http.StatusInternalServerError, Message: "internal server error"}
}

func Validation(fields []FieldError) AppError {
	return AppError{Code: "VALIDATION_ERROR", Status: http.StatusBadRequest, Message: "validation failed", Fields: fields}
}
