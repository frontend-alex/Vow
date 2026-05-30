// Package http provides shared HTTP response helpers.
package http

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"

	apperrors "vow/server/internal/shared/errors"
)

// JSON writes a JSON response with the provided status code.
func JSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if data == nil {
		return
	}

	_ = json.NewEncoder(w).Encode(data)
}

// NoContent writes a 204 response.
func NoContent(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
}

// Error writes an application error response. Unknown errors are converted to a
// generic internal error so implementation details are not leaked to clients.
func Error(w http.ResponseWriter, err error) {
	appErr := apperrors.As(err)
	JSON(w, appErr.StatusCode, ErrorBody(appErr))
}

// GinError writes an application error response through Gin.
func GinError(c *gin.Context, err error) {
	appErr := apperrors.As(err)
	c.JSON(appErr.StatusCode, ErrorBody(appErr))
}

// ErrorBody converts an application error to the standard API response body.
func ErrorBody(appErr *apperrors.Error) map[string]any {
	response := map[string]any{
		"success":     false,
		"message":     appErr.Message,
		"errorCode":   appErr.ErrorCode,
		"statusCode":  appErr.StatusCode,
		"userMessage": appErr.UserMessage,
	}

	for key, value := range appErr.Extra {
		response[key] = value
	}

	return response
}
