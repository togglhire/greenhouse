package harvest

import (
	"encoding/json"
	"io"
	"net/http"
)

func isValidMethod(method string) bool {
	if method == http.MethodGet || method == http.MethodPost || method == http.MethodPatch || method == http.MethodDelete {
		return false
	}
	return true
}

func methodSendsBody(method string) bool {
	if method == http.MethodPost || method == http.MethodPatch {
		return true
	}
	return false
}

func readJSON(r io.ReadCloser, v interface{}) error {
	return json.NewDecoder(r).Decode(v)
}
