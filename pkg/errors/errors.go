package errors

import (
	"fmt"
	"net/http"
)

// AppError represents an application error with context
type AppError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Details string `json:"details,omitempty"`
	Err     error  `json:"-"`
}

// Error implements the error interface
func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Err)
	}
	return e.Message
}

// Unwrap returns the underlying error
func (e *AppError) Unwrap() error {
	return e.Err
}

// New creates a new AppError
func New(code int, message string) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
	}
}

// Wrap wraps an existing error with context
func Wrap(err error, code int, message string) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
		Err:     err,
	}
}

// WrapWithDetails wraps an error with additional details
func WrapWithDetails(err error, code int, message, details string) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
		Details: details,
		Err:     err,
	}
}

// Predefined errors
var (
	ErrInternalServer = New(http.StatusInternalServerError, "Internal server error")
	ErrBadRequest     = New(http.StatusBadRequest, "Bad request")
	ErrUnauthorized   = New(http.StatusUnauthorized, "Unauthorized")
	ErrForbidden      = New(http.StatusForbidden, "Forbidden")
	ErrNotFound       = New(http.StatusNotFound, "Not found")
	ErrTooManyReqs    = New(http.StatusTooManyRequests, "Too many requests")
)

// API-specific errors
var (
	ErrInvalidAPIKey    = New(http.StatusUnauthorized, "Invalid API key")
	ErrAPIQuotaExceeded = New(http.StatusTooManyRequests, "API quota exceeded")
	ErrAPITimeout       = New(http.StatusRequestTimeout, "API request timeout")
	ErrInvalidRequest   = New(http.StatusBadRequest, "Invalid request format")
)

// Service-specific errors
var (
	ErrNewsServiceUnavailable   = New(http.StatusServiceUnavailable, "News service unavailable")
	ErrGeminiServiceUnavailable = New(http.StatusServiceUnavailable, "Gemini AI service unavailable")
	ErrAnalysisTimeout          = New(http.StatusRequestTimeout, "AI analysis timeout")
)
