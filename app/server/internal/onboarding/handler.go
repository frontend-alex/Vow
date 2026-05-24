package onboarding

import (
	"errors"
	"net/http"

	"github.com/vow/app/server/internal/middleware"
	sharederrors "github.com/vow/app/server/internal/shared/errors"
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
	userID, ok := userIDFromRequest(r)
	if !ok {
		response.AppError(w, sharederrors.OnboardingErrors.MissingOrInvalidUserID)
		return
	}

	onboarding, err := h.service.Start(r.Context(), userID)
	if err != nil {
		if apiError, ok := sharederrors.FromError(err); ok {
			response.AppError(w, apiError)
			return
		}

		response.AppError(w, sharederrors.OnboardingErrors.StartFailed)
		return
	}

	response.OK(w, "onboarding started", onboarding)
}

func (h Handler) Complete(w http.ResponseWriter, r *http.Request) {
	userID, ok := userIDFromRequest(r)
	if !ok {
		response.AppError(w, sharederrors.OnboardingErrors.MissingOrInvalidUserID)
		return
	}

	input, err := request.DecodeAndValidate[CompleteOnboardingRequest](w, r)
	if err != nil {
		handleRequestError(w, err)
		return
	}

	if err := h.service.Complete(r.Context(), userID, input); err != nil {
		if apiError, ok := sharederrors.FromError(err); ok {
			response.AppError(w, apiError)
			return
		}

		response.AppError(w, sharederrors.OnboardingErrors.CompleteFailed)
		return
	}

	response.OK(w, "onboarding completed", nil)
}

func handleRequestError(w http.ResponseWriter, err error) {
	var validationErr request.ValidationError
	if errors.As(err, &validationErr) {
		response.AppErrorWithMessage(w, sharederrors.RequestErrors.ValidationFailed, validationErr.Error())
		return
	}

	if errors.Is(err, request.ErrInvalidJSON) {
		response.AppError(w, sharederrors.RequestErrors.InvalidRequestBody)
		return
	}

	response.AppError(w, sharederrors.RequestErrors.InvalidRequest)
}

func userIDFromRequest(r *http.Request) (int64, bool) {
	userID, ok := middleware.UserIDFromContext(r.Context())
	if !ok {
		return 0, false
	}

	return userID, true
}
