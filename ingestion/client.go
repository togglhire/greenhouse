package ingestion

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"net/http"
	"net/url"
	"strings"
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

// Params are used to send parameters with the request.
type Params map[string]interface{}

// newRequest creates an authenticated API request that is ready to send.
func (c *Client) newRequest(method string, endpoint string, params Params, body interface{}) (*http.Request, error) {
	method = strings.ToUpper(method)
	requestURL := fmt.Sprintf("%sv1/%s", c.BaseURL, endpoint)

	// Query String
	qs := url.Values{}
	for k, v := range params {
		qs.Add(k, fmt.Sprintf("%v", v))
	}

	if len(qs) > 0 {
		requestURL += "?" + qs.Encode()
	}

	// Request body
	var buf bytes.Buffer
	if body != nil {
		err := xml.NewEncoder(&buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, requestURL, &buf)

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
	if req.Method == "POST" || req.Method == "PUT" {
		req.Header.Set("Content-Type", "application/json; charset=utf-8")
	}

	return req, err
}
