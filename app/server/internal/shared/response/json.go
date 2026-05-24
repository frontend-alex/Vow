package response

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/vow/app/server/internal/shared/apperror"
)

type APIResponse struct {
	Success bool       `json:"success"`
	Message *string    `json:"message"`
	Data    interface{} `json:"data"`
	Errors  []APIError `json:"errors"`
}

type APIError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Field   string `json:"field,omitempty"`
}

func JSON(w http.ResponseWriter, status int, payload APIResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(payload); err != nil {
		http.Error(w, `{"success":false,"message":null,"data":null,"errors":[{"code":"INTERNAL_SERVER_ERROR","message":"failed to encode response"}]}`, http.StatusInternalServerError)
	}
}

func Success(w http.ResponseWriter, status int, message string, data interface{}) {
	JSON(w, status, APIResponse{
		Success: true,
		Message: stringPtr(message),
		Data:    data,
		Errors:  []APIError{},
	})
}

func SuccessNoMessage(w http.ResponseWriter, status int, data interface{}) {
	JSON(w, status, APIResponse{
		Success: true,
		Message: nil,
		Data:    data,
		Errors:  []APIError{},
	})
}

func Error(w http.ResponseWriter, err error) {
	HandleError(w, err)
}

func HandleError(w http.ResponseWriter, err error) {
	var appErr apperror.AppError
	if errors.As(err, &appErr) {
		JSON(w, appErr.Status, APIResponse{
			Success: false,
			Message: nil,
			Data:    nil,
			Errors:  mapAppError(appErr),
		})
		return
	}

	internal := apperror.Internal()
	JSON(w, internal.Status, APIResponse{
		Success: false,
		Message: nil,
		Data:    nil,
		Errors:  mapAppError(internal),
	})
}

func OK(w http.ResponseWriter, message string, data interface{}) {
	Success(w, http.StatusOK, message, data)
}

func OKNoMessage(w http.ResponseWriter, data interface{}) {
	SuccessNoMessage(w, http.StatusOK, data)
}

func Created(w http.ResponseWriter, message string, data interface{}) {
	Success(w, http.StatusCreated, message, data)
}

func CreatedNoMessage(w http.ResponseWriter, data interface{}) {
	SuccessNoMessage(w, http.StatusCreated, data)
}

func BadRequest(w http.ResponseWriter, message string) {
	HandleError(w, apperror.BadRequest("BAD_REQUEST", message))
}

func Unauthorized(w http.ResponseWriter, message string) {
	HandleError(w, apperror.Unauthorized("UNAUTHORIZED", message))
}

func Forbidden(w http.ResponseWriter, message string) {
	HandleError(w, apperror.Forbidden("FORBIDDEN", message))
}

func NotFound(w http.ResponseWriter, message string) {
	HandleError(w, apperror.NotFound("NOT_FOUND", message))
}

func Conflict(w http.ResponseWriter, message string) {
	HandleError(w, apperror.Conflict("CONFLICT", message))
}

func InternalServerError(w http.ResponseWriter) {
	HandleError(w, apperror.Internal())
}

func mapAppError(err apperror.AppError) []APIError {
	if len(err.Fields) == 0 {
		return []APIError{{Code: err.Code, Message: err.Message}}
	}

	apiErrors := make([]APIError, 0, len(err.Fields))
	for _, field := range err.Fields {
		apiErrors = append(apiErrors, APIError{
			Code:    err.Code,
			Message: field.Message,
			Field:   field.Field,
		})
	}

	return apiErrors
}

func stringPtr(value string) *string {
	return &value
}
