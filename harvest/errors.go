package harvest

import (
	"fmt"
	"net/http"
)

type harvestError struct {
	Code    int
	Message string
}

func (e harvestError) Error() string {
	return fmt.Sprintf("Harvest Error %d: %s", e.Code, e.Message)
}

type AuthError struct {
	*harvestError
}

func NewAuthError(message string) error {
	return &AuthError{&harvestError{Code: 401, Message: message}}
}

func (ae *AuthError) Error() string {
	return fmt.Sprintf("Harvest Auth Error %d: %s", ae.Code, ae.Message)
}

type ForbiddenError struct {
	*harvestError
}

func NewForbiddenError(message string) error {
	return &ForbiddenError{&harvestError{Code: 403, Message: message}}
}

func (fe *ForbiddenError) Error() string {
	return fmt.Sprintf("Harvest Forbidden Error %d: %s", fe.Code, fe.Message)
}

type ValidationFieldError struct {
	Message string `json:"message"`
	Field   string `json:"field"`
}

type ValidationError struct {
	*harvestError
	FieldErrors []ValidationFieldError `json:"errors"`
}

func NewValidationError(message string, fieldErrors []ValidationFieldError) error {
	return &ValidationError{&harvestError{Code: 422, Message: message}, fieldErrors}
}

func (ve *ValidationError) Error() string {
	return fmt.Sprintf("Harvest Validation Error %d: %s", ve.Code, ve.Message)
}

type RateLimitError struct {
	*harvestError
}

func NewRateLimitError(message string) error {
	return &RateLimitError{&harvestError{Code: 429, Message: message}}
}

func (rle *RateLimitError) Error() string {
	return fmt.Sprintf("Harvest Rate Limit Error %d: %s", rle.Code, rle.Message)
}

type ServerError struct {
	*harvestError
}

func NewServerError(message string) error {
	return &ServerError{&harvestError{Code: 500, Message: message}}
}

func (se *ServerError) Error() string {
	return fmt.Sprintf("Harvest Server Error %d: %s", se.Code, se.Message)
}

// SDKError is a wrapper for errors that occur in the SDK itself.
type SDKError struct {
	*harvestError
}

func NewSDKError(message string) error {
	return &SDKError{&harvestError{Code: 500, Message: message}}
}

func (se *SDKError) Error() string {
	return fmt.Sprintf("Harvest SDK Error %d: %s", se.Code, se.Message)
}

func isErrorResponse(resp *http.Response) (bool, error) {
	if resp == nil {
		return false, NewSDKError("could not get a response, response is nil")
	}
	respCode := resp.StatusCode
	if respCode == 401 {
		return true, NewAuthError("authetication failed")
	}
	if respCode == 403 {
		return true, NewForbiddenError("forbidden")
	}
	if respCode == 422 {
		error := NewValidationError("invalid input provided", nil)
		if err := readJSON(resp.Body, error); err != nil {
			return false, NewSDKError(fmt.Sprintf("error decoding response: %v", err))
		}
		return true, error
	}
	if respCode == 429 {
		return true, NewRateLimitError("rate limit exceeded")
	}
	if respCode >= 500 && respCode <= 599 {
		return true, NewServerError("internal server error")
	}

	return false, nil
}
