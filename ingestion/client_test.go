package ingestion

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	// mux is the HTTP request multiplexer used with the test server
	mux *http.ServeMux

	// server is a test HTTP server used to provide mock API responses
	server *httptest.Server

	// client is the Recurly client being tested
	client *Client
)

// setup sets up a test HTTP server along with a ingestion.Client that is
// configured to talk to that test server. Tests should register handlers on
// mux which provide mock responses for the API method being tested
func setup() {
	// test server
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)

	client = NewClient("access_token", nil)
	client.BaseURL = server.URL + "/"
}

func teardown() {
	server.Close()
}

func formatReadCloser(r *io.ReadCloser) string {
	if r == nil {
		return ""
	}
	body, err := ioutil.ReadAll(*r)
	if err != nil {
		return ""
	}
	rdr1 := ioutil.NopCloser(bytes.NewBuffer(body))
	*r = rdr1 // restore body

	return string(body)
}

func areEqualJSON(s1, s2 string) (bool, error) {
	var o1 interface{}
	var o2 interface{}

	var err error
	err = json.Unmarshal([]byte(s1), &o1)
	if err != nil {
		return false, fmt.Errorf("Error mashalling string 1 :: %s", err.Error())
	}
	err = json.Unmarshal([]byte(s2), &o2)
	if err != nil {
		return false, fmt.Errorf("Error mashalling string 2 :: %s", err.Error())
	}

	return reflect.DeepEqual(o1, o2), nil
}

func Test_int64ArrayToCSV(t *testing.T) {
	type args struct {
		a []int64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "A",
			args: args{
				a: []int64{2, 3, 4, 5},
			},
			want: "2,3,4,5",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, int64ArrayToCSV(tt.args.a))
		})
	}
}
