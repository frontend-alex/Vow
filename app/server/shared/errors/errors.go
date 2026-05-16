package errors

import (
	"errors"
	"net/http"
)

type Definition struct {
	ErrorCode   string `json:"errorCode"`
	StatusCode  int    `json:"statusCode"`
	Message     string `json:"message"`
	UserMessage string `json:"userMessage"`
}

type AppError struct {
	Definition Definition
	Cause      error
}

func (e AppError) Error() string {
	if e.Cause != nil {
		return e.Definition.Message + ": " + e.Cause.Error()
	}
	return e.Definition.Message
}

func (e AppError) Unwrap() error {
	return e.Cause
}

func New(def Definition) error {
	return AppError{Definition: def}
}

func Wrap(def Definition, cause error) error {
	return AppError{Definition: def, Cause: cause}
}

var (
	AuthUnauthorized = Definition{ErrorCode: "AUTH_001", StatusCode: http.StatusUnauthorized, Message: "Unauthorized access.", UserMessage: "You must be logged in to perform this action."}
	AuthForbiddenAccess = Definition{ErrorCode: "AUTH_002", StatusCode: http.StatusForbidden, Message: "Forbidden access.", UserMessage: "You do not have permission to access this resource."}
	AuthInvalidCredentials = Definition{ErrorCode: "AUTH_003", StatusCode: http.StatusUnauthorized, Message: "Invalid email or password.", UserMessage: "The login information you provided is incorrect."}
	AuthInvalidCurrentPassword = Definition{ErrorCode: "AUTH_004", StatusCode: http.StatusBadRequest, Message: "Current password is incorrect.", UserMessage: "The current password you entered is incorrect."}
	AuthSamePassword = Definition{ErrorCode: "AUTH_005", StatusCode: http.StatusBadRequest, Message: "New password cannot be the same as the current password.", UserMessage: "Please choose a different password."}
	AuthEmailNotVerified = Definition{ErrorCode: "AUTH_006", StatusCode: http.StatusForbidden, Message: "Email has not been verified.", UserMessage: "Please verify your email before continuing."}
	AuthEmailNotProvided = Definition{ErrorCode: "AUTH_007", StatusCode: http.StatusBadRequest, Message: "Email is required.", UserMessage: "Please provide your email address."}
	AuthPasswordMissing = Definition{ErrorCode: "AUTH_008", StatusCode: http.StatusBadRequest, Message: "Password is required.", UserMessage: "Please enter your password."}
	AuthEmailAlreadyVerified = Definition{ErrorCode: "AUTH_009", StatusCode: http.StatusBadRequest, Message: "Email is already verified.", UserMessage: "Your email is already verified."}
	AuthRegistrationFailed = Definition{ErrorCode: "AUTH_010", StatusCode: http.StatusInternalServerError, Message: "Failed to register user.", UserMessage: "We couldn't complete your registration. Please try again."}
	AuthLoginFailed = Definition{ErrorCode: "AUTH_011", StatusCode: http.StatusUnauthorized, Message: "Login failed due to invalid credentials or other issues.", UserMessage: "Login failed. Please check your email and password and try again."}
	AuthUserAlreadyExists = Definition{ErrorCode: "AUTH_012", StatusCode: http.StatusBadRequest, Message: "User with this email or username already exists.", UserMessage: "An account with this email or username already exists."}
	AuthUsernameAlreadyTaken = Definition{ErrorCode: "AUTH_013", StatusCode: http.StatusBadRequest, Message: "Username is already in use.", UserMessage: "This username is already taken. Please choose another."}
	AuthPasswordResetFailed = Definition{ErrorCode: "AUTH_014", StatusCode: http.StatusInternalServerError, Message: "Failed to reset password.", UserMessage: "We couldn't reset your password. Please try again later."}
	AuthAccountLocked = Definition{ErrorCode: "AUTH_015", StatusCode: http.StatusForbidden, Message: "Account is locked due to multiple failed login attempts.", UserMessage: "Your account has been locked. Please try again later or contact support."}
	AuthEmailAlreadyTaken = Definition{ErrorCode: "AUTH_016", StatusCode: http.StatusBadRequest, Message: "Email is already in use.", UserMessage: "This email is already taken. Please use another."}
	AuthEmailNotRegistered = Definition{ErrorCode: "AUTH_017", StatusCode: http.StatusNotFound, Message: "Email is not registered.", UserMessage: "We couldn't find an account with this email address."}
	AuthAccountAlreadyConnectedWithProvider = Definition{ErrorCode: "AUTH_018", StatusCode: http.StatusBadRequest, Message: "Account already registered using a third-party provider.", UserMessage: "This email is already linked to a social login. Please sign in using that provider instead."}

	JWTInvalidToken = Definition{ErrorCode: "JWT_001", StatusCode: http.StatusUnauthorized, Message: "Invalid or expired token.", UserMessage: "Your session has expired. Please log in again."}
	JWTInvalidRefreshToken = Definition{ErrorCode: "JWT_002", StatusCode: http.StatusUnauthorized, Message: "Invalid or expired refresh token.", UserMessage: "Your session has expired. Please log in again."}
	JWTRefreshFailed = Definition{ErrorCode: "JWT_003", StatusCode: http.StatusForbidden, Message: "Token refresh failed.", UserMessage: "We couldn't refresh your session. Please log in again."}
	JWTInvalidEncryptedToken = Definition{ErrorCode: "JWT_004", StatusCode: http.StatusBadRequest, Message: "Invalid encrypted token.", UserMessage: "Authentication failed. Please try logging in again."}

	OTPExpired = Definition{ErrorCode: "OTP_001", StatusCode: http.StatusBadRequest, Message: "OTP has expired.", UserMessage: "Your OTP has expired. Please request a new one."}
	OTPInvalid = Definition{ErrorCode: "OTP_002", StatusCode: http.StatusBadRequest, Message: "Invalid OTP code.", UserMessage: "The OTP code you entered is invalid."}
	OTPNotFound = Definition{ErrorCode: "OTP_003", StatusCode: http.StatusNotFound, Message: "OTP not found.", UserMessage: "No OTP was found for your request. Please try again."}
	OTPSendFailed = Definition{ErrorCode: "OTP_004", StatusCode: http.StatusInternalServerError, Message: "Failed to send OTP.", UserMessage: "We couldn't send the OTP. Please try again later."}
	OTPAlreadyUsed = Definition{ErrorCode: "OTP_005", StatusCode: http.StatusBadRequest, Message: "OTP has already been used.", UserMessage: "This OTP has already been used. Please request a new one."}
	OTPCleanupFailed = Definition{ErrorCode: "OTP_006", StatusCode: http.StatusInternalServerError, Message: "Failed to cleanup expired OTPs.", UserMessage: "Failed to cleanup expired OTPs."}

	StripePaymentFailed = Definition{ErrorCode: "STRIPE_001", StatusCode: http.StatusPaymentRequired, Message: "Payment processing failed.", UserMessage: "Your payment could not be processed. Please try again."}
	StripeInvoiceCreationFailed = Definition{ErrorCode: "STRIPE_002", StatusCode: http.StatusInternalServerError, Message: "Invoice creation failed.", UserMessage: "We couldn't create your invoice. Please try again."}
	StripeSubscriptionCreationFailed = Definition{ErrorCode: "STRIPE_003", StatusCode: http.StatusInternalServerError, Message: "Subscription creation failed.", UserMessage: "Failed to create your subscription. Please try again."}
	StripeInvalidPlanID = Definition{ErrorCode: "STRIPE_004", StatusCode: http.StatusBadRequest, Message: "Invalid Stripe plan ID.", UserMessage: "The selected plan is not valid. Please choose another."}
	StripeCustomerCreationFailed = Definition{ErrorCode: "STRIPE_005", StatusCode: http.StatusInternalServerError, Message: "Stripe customer creation failed.", UserMessage: "We couldn't create your customer profile. Please try again."}
	StripeSubscriptionUpgradeRequired = Definition{ErrorCode: "STRIPE_006", StatusCode: http.StatusForbidden, Message: "Upgrade required for this feature.", UserMessage: "Please upgrade your plan to access this feature."}

	GoogleAuthFailed = Definition{ErrorCode: "GOOGLE_001", StatusCode: http.StatusUnauthorized, Message: "Google authentication failed.", UserMessage: "We couldn't sign you in with Google. Try again or use another method."}

	UserNotFound = Definition{ErrorCode: "USER_001", StatusCode: http.StatusNotFound, Message: "User not found.", UserMessage: "We couldn't find a user with that information."}
	UserNoUpdatesProvided = Definition{ErrorCode: "USER_002", StatusCode: http.StatusBadRequest, Message: "No update fields provided.", UserMessage: "Please provide at least one field to update."}

	EmailSendingFailed = Definition{ErrorCode: "EMAIL_001", StatusCode: http.StatusInternalServerError, Message: "Failed to send email.", UserMessage: "We couldn't send the email. Please try again later."}
	DatabaseError = Definition{ErrorCode: "DB_001", StatusCode: http.StatusInternalServerError, Message: "A database error occurred.", UserMessage: "There was a problem accessing the database. Try again later."}
	NotificationCreationFailed = Definition{ErrorCode: "NOTIF_001", StatusCode: http.StatusInternalServerError, Message: "Failed to create notification.", UserMessage: "We couldn't create the notification. Please try again."}

	BadRequest = Definition{ErrorCode: "REQ_001", StatusCode: http.StatusBadRequest, Message: "Bad request.", UserMessage: "The request is invalid."}
	Conflict = Definition{ErrorCode: "REQ_002", StatusCode: http.StatusConflict, Message: "Resource conflict.", UserMessage: "The request conflicts with the current resource state."}
	RateLimited = Definition{ErrorCode: "REQ_003", StatusCode: http.StatusTooManyRequests, Message: "Rate limit exceeded.", UserMessage: "Too many requests. Please try again later."}
	Internal = Definition{ErrorCode: "SYS_001", StatusCode: http.StatusInternalServerError, Message: "Internal server error.", UserMessage: "Something went wrong. Please try again later."}
)

var ErrorMessages = map[string]Definition{
	"AUTH_001": AuthUnauthorized,
	"AUTH_002": AuthForbiddenAccess,
	"AUTH_003": AuthInvalidCredentials,
	"AUTH_004": AuthInvalidCurrentPassword,
	"AUTH_005": AuthSamePassword,
	"AUTH_006": AuthEmailNotVerified,
	"AUTH_007": AuthEmailNotProvided,
	"AUTH_008": AuthPasswordMissing,
	"AUTH_009": AuthEmailAlreadyVerified,
	"AUTH_010": AuthRegistrationFailed,
	"AUTH_011": AuthLoginFailed,
	"AUTH_012": AuthUserAlreadyExists,
	"AUTH_013": AuthUsernameAlreadyTaken,
	"AUTH_014": AuthPasswordResetFailed,
	"AUTH_015": AuthAccountLocked,
	"AUTH_016": AuthEmailAlreadyTaken,
	"AUTH_017": AuthEmailNotRegistered,
	"AUTH_018": AuthAccountAlreadyConnectedWithProvider,
	"JWT_001": JWTInvalidToken,
	"JWT_002": JWTInvalidRefreshToken,
	"JWT_003": JWTRefreshFailed,
	"JWT_004": JWTInvalidEncryptedToken,
	"OTP_001": OTPExpired,
	"OTP_002": OTPInvalid,
	"OTP_003": OTPNotFound,
	"OTP_004": OTPSendFailed,
	"OTP_005": OTPAlreadyUsed,
	"OTP_006": OTPCleanupFailed,
	"STRIPE_001": StripePaymentFailed,
	"STRIPE_002": StripeInvoiceCreationFailed,
	"STRIPE_003": StripeSubscriptionCreationFailed,
	"STRIPE_004": StripeInvalidPlanID,
	"STRIPE_005": StripeCustomerCreationFailed,
	"STRIPE_006": StripeSubscriptionUpgradeRequired,
	"GOOGLE_001": GoogleAuthFailed,
	"USER_001": UserNotFound,
	"USER_002": UserNoUpdatesProvided,
	"EMAIL_001": EmailSendingFailed,
	"DB_001": DatabaseError,
	"NOTIF_001": NotificationCreationFailed,
	"REQ_001": BadRequest,
	"REQ_002": Conflict,
	"REQ_003": RateLimited,
	"SYS_001": Internal,
}

var (
	ErrBadRequest = New(BadRequest)
	ErrUnauthorized = New(AuthUnauthorized)
	ErrForbidden = New(AuthForbiddenAccess)
	ErrNotFound = New(UserNotFound)
	ErrConflict = New(Conflict)
	ErrRateLimited = New(RateLimited)
	ErrInternal = New(Internal)
)

func Details(err error) Definition {
	var appErr AppError
	if errors.As(err, &appErr) {
		return appErr.Definition
	}

	return Internal
}

func HTTPError(err error) (int, string) {
	details := Details(err)
	return details.StatusCode, details.UserMessage
}
