package response

import (
	"encoding/json"
	"net/http"

	apperrors "github.com/frontend-alex/Vow/app/server/shared/errors"
)

type envelope map[string]any

func JSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(envelope{
		"success":    status >= 200 && status < 300,
		"statusCode": status,
		"data":       data,
	})
}

func Error(w http.ResponseWriter, r *http.Request, err error) {
	details := apperrors.Details(err)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(details.StatusCode)
	_ = json.NewEncoder(w).Encode(envelope{
		"success":    false,
		"statusCode": details.StatusCode,
		"error": envelope{
			"errorCode":   details.ErrorCode,
			"message":     details.Message,
			"userMessage": details.UserMessage,
			"requestId":   r.Header.Get("X-Request-ID"),
		},
	})
}
