package harvest

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

const (
	TEST_API_KEY      = "api_key"
	TEST_ON_BEHALF_OF = "harvest-sdk-client"
)

var (
	//  Multiplexer used with the test server
	mux *http.ServeMux
	// Test HTTP server used to provide mock API responses
	server *httptest.Server

	// Client being tested
	client *Client
)

func setup(apiKey string, onBehalfOf string, t *testing.T) {
	// Test server
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)

	var err error
	client, err = NewDefaultClient(apiKey, onBehalfOf, nil)
	if err != nil {
		t.Errorf("Error creating client: %s", err.Error())
	}
}

func teardown() {
	server.Close()
}

func TestNewDefaultClient(t *testing.T) {
	setup(TEST_API_KEY, TEST_ON_BEHALF_OF, t)
	defer teardown()

	if client.baseURL != DEFAULT_BASE_URL+string(V1)+"/" {
		t.Errorf("NewDefaultClient baseURL: %s, expected %s", client.baseURL, DEFAULT_BASE_URL+string(V1)+"/")
	}

	if client.apiKey != TEST_API_KEY {
		t.Errorf("NewDefaultClient apiKey: %s, expected: %s", client.apiKey, TEST_API_KEY)
	}

	if client.onBehalfOf != TEST_ON_BEHALF_OF {
		t.Errorf("NewDefaultClient onBehalfOf: %s, expected: %s", client.onBehalfOf, TEST_ON_BEHALF_OF)
	}
}

func TestNewDefaultClientInvalidAPIKey(t *testing.T) {
	_, err := NewDefaultClient("", "", nil)
	if err == nil {
		t.Errorf("Expected error creating client with an empty api key")
	}

	var sdkError *SDKError
	if !errors.As(err, &sdkError) {
		t.Errorf("Expected error to be AuthError")
	}
}

func TestClientNewRequestWithInvalidMethod(t *testing.T) {
	setup(TEST_API_KEY, TEST_ON_BEHALF_OF, t)
	defer teardown()

	_, err := client.newRequest("INVALID", "test", nil, nil)
	if err == nil {
		t.Errorf("Expected error creating request with invalid method")
	}

	var sdkError *SDKError
	if !errors.As(err, &sdkError) {
		t.Errorf("Expected error to be AuthError")
	}
}
