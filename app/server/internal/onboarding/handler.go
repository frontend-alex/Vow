package onboarding

import (
	"errors"
	"net/http"

	"github.com/vow/app/server/internal/middleware"
	"github.com/vow/app/server/internal/shared/request"
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
		if errors.Is(err, ErrOnboardingAlreadyStarted) {
			response.Error(w, http.StatusConflict, "onboarding already started")
			return
		}

		if errors.Is(err, ErrOnboardingAlreadyCompleted) {
			response.Error(w, http.StatusConflict, "onboarding already completed")
			return
		}

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

	input, err := request.DecodeAndValidate[CompleteOnboardingRequest](w, r)
	if err != nil {
		handleRequestError(w, err)
		return
	}

	if err := h.service.Complete(r.Context(), userID, input); err != nil {
		if errors.Is(err, ErrOnboardingAlreadyCompleted) {
			response.Error(w, http.StatusConflict, "onboarding already completed")
			return
		}

		if errors.Is(err, ErrOnboardingNotStarted) {
			response.Error(w, http.StatusBadRequest, "onboarding not started")
			return
		}

		response.Error(w, http.StatusInternalServerError, "failed to complete onboarding")
		return
	}

	response.OK(w, "onboarding completed", nil)
}

func handleRequestError(w http.ResponseWriter, err error) {
	var validationErr request.ValidationError
	if errors.As(err, &validationErr) {
		response.Error(w, http.StatusBadRequest, validationErr.Error())
		return
	}

	if errors.Is(err, request.ErrInvalidJSON) {
		response.Error(w, http.StatusBadRequest, "invalid request body")
		return
	}

	response.Error(w, http.StatusBadRequest, "invalid request")
}

func userIDFromRequest(r *http.Request) (int64, error) {
	userID, ok := middleware.UserIDFromContext(r.Context())
	if !ok {
		return 0, errors.New("missing user id")
	}

	return userID, nil
}
