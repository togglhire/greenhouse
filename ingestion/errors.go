package ingestion

import (
	"errors"
	"fmt"
	"net/http"
)

// ErrNotImplemented is used for returning errors
var ErrNotImplemented = errors.New("Not implemented")
var ErrShouldNotBeNil = errors.New("Should not be nil")

type Error struct {
	Message string `json:"message"`
	Field   string `json:"field"`
}
type ClientError struct {
	Errors []Error `json:"errors"`
}

func (e ClientError) Error() string {
	return fmt.Sprintf("var: %#+v\n", e.Errors)
}

type ServerError struct {
	Errors []Error `json:"errors"`
}

func (e ServerError) Error() string {
	return fmt.Sprintf("var: %#+v\n", e.Errors)
}

func isOK(r *http.Response) (bool, error) {
	if r == nil {
		return false, ErrShouldNotBeNil
	}
	return r.StatusCode >= 200 && r.StatusCode <= 299, nil
}

func isError(r *http.Response) (bool, error) {
	if r == nil {
		return false, ErrShouldNotBeNil
	}
	ok, err := isOK(r)
	return !ok, err
}

func isClientError(r *http.Response) (bool, error) {
	if r == nil {
		return false, ErrShouldNotBeNil
	}
	return r.StatusCode >= 400 && r.StatusCode <= 499, nil
}

func isServerError(r *http.Response) (bool, error) {
	if r == nil {
		return false, ErrShouldNotBeNil
	}
	return r.StatusCode >= 500 && r.StatusCode <= 599, nil
}
