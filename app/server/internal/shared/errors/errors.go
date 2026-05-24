package errors

import (
	stderrors "errors"
	"net/http"
)

type APIError struct {
	ErrorCode   string
	StatusCode  int
	Message     string
	UserMessage string
}

func (e APIError) Error() string {
	return e.Message
}

func FromError(err error) (APIError, bool) {
	var apiError APIError
	if stderrors.As(err, &apiError) {
		return apiError, true
	}

	return APIError{}, false
}

var AuthErrors = struct {
	Unauthorized       APIError
	ForbiddenAccess    APIError
	InvalidCredentials APIError
	EmailAlreadyTaken  APIError
}{
	Unauthorized: APIError{
		ErrorCode:   "AUTH_001",
		StatusCode:  http.StatusUnauthorized,
		Message:     "Unauthorized access.",
		UserMessage: "You must be logged in to perform this action.",
	},
	ForbiddenAccess: APIError{
		ErrorCode:   "AUTH_002",
		StatusCode:  http.StatusForbidden,
		Message:     "Forbidden access.",
		UserMessage: "You do not have permission to access this resource.",
	},
	InvalidCredentials: APIError{
		ErrorCode:   "AUTH_003",
		StatusCode:  http.StatusUnauthorized,
		Message:     "Invalid email or password.",
		UserMessage: "The login information you provided is incorrect.",
	},
	EmailAlreadyTaken: APIError{
		ErrorCode:   "AUTH_016",
		StatusCode:  http.StatusConflict,
		Message:     "Email is already in use.",
		UserMessage: "This email is already taken. Please use another.",
	},
}

var JWTErrors = struct {
	InvalidToken APIError
}{
	InvalidToken: APIError{
		ErrorCode:   "JWT_001",
		StatusCode:  http.StatusUnauthorized,
		Message:     "Invalid or expired token.",
		UserMessage: "Your session has expired. Please log in again.",
	},
}

var RequestErrors = struct {
	InvalidRequestBody APIError
	InvalidRequest     APIError
	ValidationFailed   APIError
}{
	InvalidRequestBody: APIError{
		ErrorCode:   "REQ_001",
		StatusCode:  http.StatusBadRequest,
		Message:     "Invalid request body.",
		UserMessage: "Please check your request and try again.",
	},
	InvalidRequest: APIError{
		ErrorCode:   "REQ_002",
		StatusCode:  http.StatusBadRequest,
		Message:     "Invalid request.",
		UserMessage: "Please check your request and try again.",
	},
	ValidationFailed: APIError{
		ErrorCode:   "REQ_003",
		StatusCode:  http.StatusBadRequest,
		Message:     "Validation failed.",
		UserMessage: "Please check the highlighted fields and try again.",
	},
}

var OnboardingErrors = struct {
	MissingOrInvalidUserID APIError
	AlreadyStarted         APIError
	AlreadyCompleted       APIError
	NotStarted             APIError
	StartFailed            APIError
	CompleteFailed         APIError
}{
	MissingOrInvalidUserID: APIError{
		ErrorCode:   "ONBOARDING_001",
		StatusCode:  http.StatusUnauthorized,
		Message:     "Missing or invalid user id.",
		UserMessage: "You must be logged in to perform this action.",
	},
	AlreadyStarted: APIError{
		ErrorCode:   "ONBOARDING_002",
		StatusCode:  http.StatusConflict,
		Message:     "Onboarding already started.",
		UserMessage: "Your onboarding has already been started.",
	},
	AlreadyCompleted: APIError{
		ErrorCode:   "ONBOARDING_003",
		StatusCode:  http.StatusConflict,
		Message:     "Onboarding already completed.",
		UserMessage: "Your onboarding has already been completed.",
	},
	NotStarted: APIError{
		ErrorCode:   "ONBOARDING_004",
		StatusCode:  http.StatusBadRequest,
		Message:     "Onboarding not started.",
		UserMessage: "Please start onboarding before completing it.",
	},
	StartFailed: APIError{
		ErrorCode:   "ONBOARDING_005",
		StatusCode:  http.StatusInternalServerError,
		Message:     "Failed to start onboarding.",
		UserMessage: "We couldn't start onboarding. Please try again.",
	},
	CompleteFailed: APIError{
		ErrorCode:   "ONBOARDING_006",
		StatusCode:  http.StatusInternalServerError,
		Message:     "Failed to complete onboarding.",
		UserMessage: "We couldn't complete onboarding. Please try again.",
	},
}

var GeneralErrors = struct {
	InternalServerError APIError
	NotImplemented      APIError
	TooManyRequests     APIError
}{
	InternalServerError: APIError{
		ErrorCode:   "GEN_001",
		StatusCode:  http.StatusInternalServerError,
		Message:     "Internal server error.",
		UserMessage: "Something went wrong. Please try again later.",
	},
	NotImplemented: APIError{
		ErrorCode:   "GEN_002",
		StatusCode:  http.StatusNotImplemented,
		Message:     "Not implemented.",
		UserMessage: "This feature is not available yet.",
	},
	TooManyRequests: APIError{
		ErrorCode:   "GEN_003",
		StatusCode:  http.StatusTooManyRequests,
		Message:     "Too many requests.",
		UserMessage: "You are sending too many requests. Please try again later.",
	},
}

var ErrorMessages = map[string]APIError{
	AuthErrors.Unauthorized.ErrorCode:                 AuthErrors.Unauthorized,
	AuthErrors.ForbiddenAccess.ErrorCode:              AuthErrors.ForbiddenAccess,
	AuthErrors.InvalidCredentials.ErrorCode:           AuthErrors.InvalidCredentials,
	AuthErrors.EmailAlreadyTaken.ErrorCode:            AuthErrors.EmailAlreadyTaken,
	JWTErrors.InvalidToken.ErrorCode:                  JWTErrors.InvalidToken,
	RequestErrors.InvalidRequestBody.ErrorCode:        RequestErrors.InvalidRequestBody,
	RequestErrors.InvalidRequest.ErrorCode:            RequestErrors.InvalidRequest,
	RequestErrors.ValidationFailed.ErrorCode:          RequestErrors.ValidationFailed,
	OnboardingErrors.MissingOrInvalidUserID.ErrorCode: OnboardingErrors.MissingOrInvalidUserID,
	OnboardingErrors.AlreadyStarted.ErrorCode:         OnboardingErrors.AlreadyStarted,
	OnboardingErrors.AlreadyCompleted.ErrorCode:       OnboardingErrors.AlreadyCompleted,
	OnboardingErrors.NotStarted.ErrorCode:             OnboardingErrors.NotStarted,
	OnboardingErrors.StartFailed.ErrorCode:            OnboardingErrors.StartFailed,
	OnboardingErrors.CompleteFailed.ErrorCode:         OnboardingErrors.CompleteFailed,
	GeneralErrors.InternalServerError.ErrorCode:       GeneralErrors.InternalServerError,
	GeneralErrors.NotImplemented.ErrorCode:            GeneralErrors.NotImplemented,
	GeneralErrors.TooManyRequests.ErrorCode:           GeneralErrors.TooManyRequests,
}
