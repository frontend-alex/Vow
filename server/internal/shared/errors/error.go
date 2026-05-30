// Package errors provides catalog-based application errors that can be safely
// rendered to API clients.
package errors

import (
	stderrors "errors"
	"fmt"
	"net/http"
)

type Key string

const (
	Unknown          Key = "UNKNOWN"
	ValidationFailed Key = "VALIDATION_FAILED"

	AuthUnauthorized                        Key = "AUTH_UNAUTHORIZED"
	AuthForbiddenAccess                     Key = "AUTH_FORBIDDEN_ACCESS"
	AuthInvalidCredentials                  Key = "AUTH_INVALID_CREDENTIALS"
	AuthInvalidCurrentPassword              Key = "AUTH_INVALID_CURRENT_PASSWORD"
	AuthSamePassword                        Key = "AUTH_SAME_PASSWORD"
	AuthEmailAlreadyTaken                   Key = "AUTH_EMAIL_ALREADY_TAKEN"
	AuthAccountAlreadyConnectedWithProvider Key = "AUTH_ACCOUNT_ALREADY_CONNECTED_WITH_PROVIDER"
	AuthEmailNotVerified                    Key = "AUTH_EMAIL_NOT_VERIFIED"
	AuthEmailNotProvided                    Key = "AUTH_EMAIL_NOT_PROVIDED"
	AuthPasswordMissing                     Key = "AUTH_PASSWORD_MISSING"
	AuthEmailAlreadyVerified                Key = "AUTH_EMAIL_ALREADY_VERIFIED"
	AuthRegistrationFailed                  Key = "AUTH_REGISTRATION_FAILED"
	AuthLoginFailed                         Key = "AUTH_LOGIN_FAILED"
	AuthUserAlreadyExists                   Key = "AUTH_USER_ALREADY_EXISTS"
	AuthUsernameAlreadyTaken                Key = "AUTH_USERNAME_ALREADY_TAKEN"
	AuthPasswordResetFailed                 Key = "AUTH_PASSWORD_RESET_FAILED"
	AuthAccountLocked                       Key = "AUTH_ACCOUNT_LOCKED"
	AuthEmailNotRegistered                  Key = "AUTH_EMAIL_NOT_REGISTERED"

	JWTInvalidToken          Key = "JWT_INVALID_TOKEN"
	JWTInvalidRefreshToken   Key = "JWT_INVALID_REFRESH_TOKEN"
	JWTRefreshFailed         Key = "JWT_REFRESH_FAILED"
	JWTInvalidEncryptedToken Key = "JWT_INVALID_ENCRYPTED_TOKEN"

	OTPExpired       Key = "OTP_EXPIRED"
	OTPInvalid       Key = "OTP_INVALID"
	OTPNotFound      Key = "OTP_NOT_FOUND"
	OTPSendFailed    Key = "OTP_SEND_FAILED"
	OTPAlreadyUsed   Key = "OTP_ALREADY_USED"
	OTPCleanupFailed Key = "OTP_CLEANUP_FAILED"

	GoogleAuthFailed Key = "GOOGLE_AUTH_FAILED"

	UserNotFound          Key = "USER_NOT_FOUND"
	UserNoUpdatesProvided Key = "USER_NO_UPDATES_PROVIDED"

	EmailSendingFailed Key = "EMAIL_SENDING_FAILED"
	DatabaseError      Key = "DATABASE_ERROR"

	NotificationCreationFailed Key = "NOTIFICATION_CREATION_FAILED"
)

// Definition is the reusable catalog entry for an application error.
type Definition struct {
	ErrorCode   string
	StatusCode  int
	Message     string
	UserMessage string
}

// Error is the concrete application error returned through the API.
type Error struct {
	Message     string         `json:"message"`
	ErrorCode   string         `json:"errorCode"`
	StatusCode  int            `json:"statusCode"`
	UserMessage string         `json:"userMessage"`
	Extra       map[string]any `json:"-"`
	Err         error          `json:"-"`
}

func (e *Error) Error() string {
	if e.Err == nil {
		return e.Message
	}
	return fmt.Sprintf("%s: %v", e.Message, e.Err)
}

func (e *Error) Unwrap() error {
	return e.Err
}

var catalog = map[Key]Definition{
	Unknown:          {"UNKNOWN", http.StatusInternalServerError, "Unknown server error.", "Something went wrong."},
	ValidationFailed: {"VALIDATION_001", http.StatusBadRequest, "Request validation failed.", "Please check your input and try again."},

	AuthUnauthorized:                        {"AUTH_001", http.StatusUnauthorized, "Unauthorized access.", "You must be logged in to perform this action."},
	AuthForbiddenAccess:                     {"AUTH_002", http.StatusForbidden, "Forbidden access.", "You do not have permission to access this resource."},
	AuthInvalidCredentials:                  {"AUTH_003", http.StatusUnauthorized, "Invalid email or password.", "The login information you provided is incorrect."},
	AuthInvalidCurrentPassword:              {"AUTH_004", http.StatusBadRequest, "Current password is incorrect.", "The current password you entered is incorrect."},
	AuthSamePassword:                        {"AUTH_005", http.StatusBadRequest, "New password cannot be the same as the current password.", "Please choose a different password."},
	AuthEmailNotVerified:                    {"AUTH_006", http.StatusForbidden, "Email has not been verified.", "Please verify your email before continuing."},
	AuthEmailNotProvided:                    {"AUTH_007", http.StatusBadRequest, "Email is required.", "Please provide your email address."},
	AuthPasswordMissing:                     {"AUTH_008", http.StatusBadRequest, "Password is required.", "Please enter your password."},
	AuthEmailAlreadyVerified:                {"AUTH_009", http.StatusBadRequest, "Email is already verified.", "Your email is already verified."},
	AuthRegistrationFailed:                  {"AUTH_010", http.StatusInternalServerError, "Failed to register user.", "We couldn't complete your registration. Please try again."},
	AuthLoginFailed:                         {"AUTH_011", http.StatusUnauthorized, "Login failed due to invalid credentials or other issues.", "Login failed. Please check your email and password and try again."},
	AuthUserAlreadyExists:                   {"AUTH_012", http.StatusBadRequest, "User with this email or username already exists.", "An account with this email or username already exists."},
	AuthUsernameAlreadyTaken:                {"AUTH_013", http.StatusBadRequest, "Username is already in use.", "This username is already taken. Please choose another."},
	AuthPasswordResetFailed:                 {"AUTH_014", http.StatusInternalServerError, "Failed to reset password.", "We couldn't reset your password. Please try again later."},
	AuthAccountLocked:                       {"AUTH_015", http.StatusForbidden, "Account is locked due to multiple failed login attempts.", "Your account has been locked. Please try again later or contact support."},
	AuthEmailAlreadyTaken:                   {"AUTH_016", http.StatusBadRequest, "Email is already in use.", "This email is already taken. Please use another."},
	AuthEmailNotRegistered:                  {"AUTH_017", http.StatusNotFound, "Email is not registered.", "We couldn't find an account with this email address."},
	AuthAccountAlreadyConnectedWithProvider: {"AUTH_018", http.StatusBadRequest, "Account already registered using a third-party provider.", "This email is already linked to a social login. Please sign in using that provider instead."},

	JWTInvalidToken:          {"JWT_001", http.StatusUnauthorized, "Invalid or expired token.", "Your session has expired. Please log in again."},
	JWTInvalidRefreshToken:   {"JWT_002", http.StatusUnauthorized, "Invalid or expired refresh token.", "Your session has expired. Please log in again."},
	JWTRefreshFailed:         {"JWT_003", http.StatusForbidden, "Token refresh failed.", "We couldn't refresh your session. Please log in again."},
	JWTInvalidEncryptedToken: {"JWT_004", http.StatusBadRequest, "Invalid encrypted token.", "Authentication failed. Please try logging in again."},

	OTPExpired:       {"OTP_001", http.StatusBadRequest, "OTP has expired.", "Your OTP has expired. Please request a new one."},
	OTPInvalid:       {"OTP_002", http.StatusBadRequest, "Invalid OTP code.", "The OTP code you entered is invalid."},
	OTPNotFound:      {"OTP_003", http.StatusNotFound, "OTP not found.", "No OTP was found for your request. Please try again."},
	OTPSendFailed:    {"OTP_004", http.StatusInternalServerError, "Failed to send OTP.", "We couldn't send the OTP. Please try again later."},
	OTPAlreadyUsed:   {"OTP_005", http.StatusBadRequest, "OTP has already been used.", "This OTP has already been used. Please request a new one."},
	OTPCleanupFailed: {"OTP_006", http.StatusInternalServerError, "Failed to cleanup expired OTPs.", "Failed to cleanup expired OTPs."},

	GoogleAuthFailed: {"GOOGLE_001", http.StatusUnauthorized, "Google authentication failed.", "We couldn't sign you in with Google. Try again or use another method."},

	UserNotFound:          {"USER_001", http.StatusNotFound, "User not found.", "We couldn't find a user with that information."},
	UserNoUpdatesProvided: {"USER_002", http.StatusBadRequest, "No update fields provided.", "Please provide at least one field to update."},

	EmailSendingFailed: {"EMAIL_001", http.StatusInternalServerError, "Failed to send email.", "We couldn't send the email. Please try again later."},
	DatabaseError:      {"DB_001", http.StatusInternalServerError, "A database error occurred.", "There was a problem accessing the database. Try again later."},

	NotificationCreationFailed: {"NOTIF_001", http.StatusInternalServerError, "Failed to create notification.", "We couldn't create the notification. Please try again."},
}

type Option func(*Error)

func WithCause(err error) Option {
	return func(e *Error) {
		e.Err = err
	}
}

func WithUserMessage(message string) Option {
	return func(e *Error) {
		e.UserMessage = message
	}
}

func WithExtra(extra map[string]any) Option {
	return func(e *Error) {
		e.Extra = extra
	}
}

// New creates an application error from a catalog key.
func New(key Key, opts ...Option) *Error {
	definition, ok := catalog[key]
	if !ok {
		definition = catalog[Unknown]
	}

	appErr := &Error{
		Message:     definition.Message,
		ErrorCode:   definition.ErrorCode,
		StatusCode:  definition.StatusCode,
		UserMessage: definition.UserMessage,
	}

	for _, opt := range opts {
		opt(appErr)
	}

	return appErr
}

func Internal(err error) *Error {
	return New(Unknown, WithCause(err))
}

// As returns err as an *Error. Unknown errors become a generic internal error.
func As(err error) *Error {
	if err == nil {
		return nil
	}

	var appErr *Error
	if stderrors.As(err, &appErr) {
		return appErr
	}

	return Internal(err)
}
