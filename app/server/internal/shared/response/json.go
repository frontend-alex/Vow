package response

import (
	"encoding/json"
	"net/http"
)

type APIResponse struct {
	Success      bool        `json:"success"`
	Message      *string     `json:"message"`
	ErrorMessage *string     `json:"error_message"`
	ErrorStatus  *int        `json:"error_status"`
	Data         interface{} `json:"data"`
}

func JSON(w http.ResponseWriter, status int, payload APIResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(payload); err != nil {
		http.Error(w, `{"success":false,"message":null,"error_message":"failed to encode response","error_status":500,"data":null}`, http.StatusInternalServerError)
	}
}

func Success(w http.ResponseWriter, status int, message string, data interface{}) {
	JSON(w, status, APIResponse{
		Success:      true,
		Message:      stringPtr(message),
		ErrorMessage: nil,
		ErrorStatus:  nil,
		Data:         data,
	})
}

func SuccessNoMessage(w http.ResponseWriter, status int, data interface{}) {
	JSON(w, status, APIResponse{
		Success:      true,
		Message:      nil,
		ErrorMessage: nil,
		ErrorStatus:  nil,
		Data:         data,
	})
}

func Error(w http.ResponseWriter, status int, message string) {
	JSON(w, status, APIResponse{
		Success:      false,
		Message:      nil,
		ErrorMessage: stringPtr(message),
		ErrorStatus:  intPtr(status),
		Data:         nil,
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
	Error(w, http.StatusBadRequest, message)
}

func Unauthorized(w http.ResponseWriter, message string) {
	Error(w, http.StatusUnauthorized, message)
}

func Forbidden(w http.ResponseWriter, message string) {
	Error(w, http.StatusForbidden, message)
}

func NotFound(w http.ResponseWriter, message string) {
	Error(w, http.StatusNotFound, message)
}

func Conflict(w http.ResponseWriter, message string) {
	Error(w, http.StatusConflict, message)
}

func InternalServerError(w http.ResponseWriter) {
	Error(w, http.StatusInternalServerError, "internal server error")
}

func stringPtr(value string) *string {
	return &value
}

func intPtr(value int) *int {
	return &value
}
