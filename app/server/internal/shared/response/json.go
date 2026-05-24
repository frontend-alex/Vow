package response

import (
	"encoding/json"
	"net/http"

	sharederrors "github.com/vow/app/server/internal/shared/errors"
)

type APIResponse struct {
	Success      bool        `json:"success"`
	Message      *string     `json:"message"`
	ErrorMessage *string     `json:"error_message"`
	ErrorStatus  *int        `json:"error_status"`
	ErrorCode    *string     `json:"error_code"`
	UserMessage  *string     `json:"user_message"`
	Data         interface{} `json:"data"`
}

func JSON(w http.ResponseWriter, status int, payload APIResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(payload); err != nil {
		http.Error(w, `{"success":false,"message":null,"error_message":"failed to encode response","error_status":500,"error_code":"GEN_001","user_message":"Something went wrong. Please try again later.","data":null}`, http.StatusInternalServerError)
	}
}

func Success(w http.ResponseWriter, status int, message string, data interface{}) {
	JSON(w, status, APIResponse{
		Success:      true,
		Message:      stringPtr(message),
		ErrorMessage: nil,
		ErrorStatus:  nil,
		ErrorCode:    nil,
		UserMessage:  nil,
		Data:         data,
	})
}

func SuccessNoMessage(w http.ResponseWriter, status int, data interface{}) {
	JSON(w, status, APIResponse{
		Success:      true,
		Message:      nil,
		ErrorMessage: nil,
		ErrorStatus:  nil,
		ErrorCode:    nil,
		UserMessage:  nil,
		Data:         data,
	})
}

func AppError(w http.ResponseWriter, apiError sharederrors.APIError) {
	AppErrorWithMessage(w, apiError, apiError.Message)
}

func AppErrorWithMessage(w http.ResponseWriter, apiError sharederrors.APIError, message string) {
	JSON(w, apiError.StatusCode, APIResponse{
		Success:      false,
		Message:      nil,
		ErrorMessage: stringPtr(message),
		ErrorStatus:  intPtr(apiError.StatusCode),
		ErrorCode:    stringPtr(apiError.ErrorCode),
		UserMessage:  stringPtr(apiError.UserMessage),
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

func InternalServerError(w http.ResponseWriter) {
	AppError(w, sharederrors.GeneralErrors.InternalServerError)
}

func stringPtr(value string) *string {
	return &value
}

func intPtr(value int) *int {
	return &value
}
