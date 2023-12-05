package harvest

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
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

func structToURLValues(data interface{}) url.Values {
	values := url.Values{}
	val := reflect.ValueOf(data)
	typ := reflect.TypeOf(data)

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		tag := typ.Field(i).Tag.Get("url")

		// Omit field if "omitempty" tag is present and it is empty
		tagOptions := strings.Split(tag, ",")
		if tagOptions[1] == "omitempty" && field.IsZero() {
			continue
		}

		// Convert non-zero values to string and add to url.Values map
		values.Add(tagOptions[0], convertValueToString(field.Interface()))
	}

	return values
}

func convertValueToString(value interface{}) string {
	switch v := value.(type) {
	case int64:
		return strconv.FormatInt(v, 10)
	case string:
		return v
	default:
		// Fallback to fmt.Sprintf for unsupported types
		return fmt.Sprintf("%v", v)
	}
}
