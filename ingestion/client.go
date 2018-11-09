package ingestion

import (
	"net/http"
)

const defaultBaseURL = "https://api.greenhouse.io/"

// Client manages communication with the Greenhouse API.
type Client struct {
	// client is the HTTP Client used to communicate with the API.
	client *http.Client

	// The access token you received once the OAuth process is complete and the user grants the partner permission to access their data on Greenhouse
	accessToken string

	// BaseURL is the base url for api requests.
	BaseURL string

	// Services used for talking with different parts of the Greenhouse API
}

// NewClient returns a new instance of *Client.
func NewClient(accessToken string, httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	client := &Client{
		client:      httpClient,
		accessToken: accessToken,
		BaseURL:     defaultBaseURL,
	}

	//Services

	return client
}
