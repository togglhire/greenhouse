package ingestion

import (
	"net/http"
	"net/http/httptest"
	"testing"
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
			if got := int64ArrayToCSV(tt.args.a); got != tt.want {
				t.Errorf("int64ArrayToCSV() = %v, want %v", got, tt.want)
			}
		})
	}
}
