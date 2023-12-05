package harvest

import (
	"fmt"
	"net/http"
)

type HarvestError struct {
	Code    int
	Message string
}

func (e HarvestError) Error() string {
	return fmt.Sprintf("Harvest Error %d: %s", e.Code, e.Message)
}

type AuthError struct {
	*HarvestError
}

func NewAuthError(message string) error {
	return &AuthError{&HarvestError{Code: 401, Message: message}}
}

type ForbiddenError struct {
	*HarvestError
}

func NewForbiddenError(message string) error {
	return &ForbiddenError{&HarvestError{Code: 403, Message: message}}
}

type ValidationFieldError struct {
	Message string `json:"message"`
	Field   string `json:"field"`
}

type ValidationError struct {
	*HarvestError
	FieldErrors []ValidationFieldError `json:"errors"`
}

func NewValidationError(message string, fieldErrors []ValidationFieldError) error {
	return &ValidationError{&HarvestError{Code: 422, Message: message}, fieldErrors}
}

type RateLimitError struct {
	*HarvestError
}

func NewRateLimitError(message string) error {
	return &RateLimitError{&HarvestError{Code: 429, Message: message}}
}

type ServerError struct {
	*HarvestError
}

func NewServerError(message string) error {
	return &ServerError{&HarvestError{Code: 500, Message: message}}
}

// SDKError is a wrapper for errors that occur in the SDK itself.
type SDKError struct {
	*HarvestError
}

func NewSDKError(message string) error {
	return &SDKError{&HarvestError{Code: 500, Message: message}}
}

func isErrorResponse(resp *http.Response) (bool, error) {
	if resp == nil {
		return false, NewSDKError("response should not be nil")
	}
	respCode := resp.StatusCode
	if respCode == 401 {
		return true, NewAuthError("authentication failed")
	}
	if respCode == 403 {
		return true, NewForbiddenError("forbidden")
	}
	if respCode == 422 {
		error := NewValidationError("validation failed", nil)
		if err := readJSON(resp.Body, error); err != nil {
			return false, NewSDKError(fmt.Sprintf("error decoding response: %v", err))
		}
		return true, error
	}
	if respCode == 429 {
		return true, NewRateLimitError("rate limit exceeded")
	}
	if respCode >= 500 && respCode <= 599 {
		return true, NewServerError("server error")
	}

	return false, nil
}
