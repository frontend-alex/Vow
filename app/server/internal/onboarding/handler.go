package onboarding

import (
	"encoding/json"
	"net/http"
	"strconv"

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
	// Temporary approach until auth middleware exists.
	// Use header: X-User-ID: 1
	raw := r.Header.Get("X-User-ID")
	return strconv.ParseInt(raw, 10, 64)
}
