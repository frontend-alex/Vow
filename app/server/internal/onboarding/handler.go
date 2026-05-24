package onboarding

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/vow/app/server/internal/middleware"
	"github.com/vow/app/server/internal/shared/response"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) Handler {
	return Handler{service: service}
}

func (h Handler) Start(w http.ResponseWriter, r *http.Request) {
	userID, err := userIDFromRequest(r)
	if err != nil {
		response.Error(w, http.StatusUnauthorized, "missing or invalid user id")
		return
	}

	onboarding, err := h.service.Start(r.Context(), userID)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "failed to start onboarding")
		return
	}

	response.OK(w, "onboarding started", onboarding)
}

func (h Handler) Complete(w http.ResponseWriter, r *http.Request) {
	userID, err := userIDFromRequest(r)
	if err != nil {
		response.Error(w, http.StatusUnauthorized, "missing or invalid user id")
		return
	}

	var input CompleteOnboardingRequest
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		response.Error(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if err := h.service.Complete(r.Context(), userID, input); err != nil {
		response.Error(w, http.StatusInternalServerError, "failed to complete onboarding")
		return
	}

	response.OK(w, "onboarding completed", nil)
}

func userIDFromRequest(r *http.Request) (int64, error) {
	userID, ok := middleware.UserIDFromContext(r.Context())
	if !ok {
		return 0, errors.New("missing user id")
	}

	return userID, nil
}
