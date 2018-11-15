package ingestion

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const (
	defaultBaseURL = "https://api.greenhouse.io/"
	accessTokenURL = "https://api.greenhouse.io/oauth/token"
	authorizeURL   = "https://api.greenhouse.io/oauth/authorize"
)

// Client manages communication with the Greenhouse API.
type Client struct {
	// client is the HTTP Client used to communicate with the API.
	client *http.Client

	// OAuth, The access token you received once the OAuth process is complete and the user grants the partner permission to access their data on Greenhouse
	accessToken string

	// Basic Auth
	apiKey     string
	onBehalfOf string

	// BaseURL is the base url for api requests.
	baseURL string

	// Services used for talking with different parts of the Greenhouse API
	OAuth         OAuthService
	Candidates    CandidateService
	CurrentUser   CurrentUserService
	Jobs          JobService
	TrackingLinks TrackingLinkService
}

// NewClient returns a new instance of *Client.
func NewClient(accessToken string, httpClient *http.Client) *Client {
	return newClient(accessToken, "", "", httpClient)
}

func NewClientBasicAuth(apiKey string, onBehalfOf string, httpClient *http.Client) *Client {
	return newClient("", apiKey, onBehalfOf, httpClient)
}

func newClient(accessToken, apiKey, onBehalfOf string, httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	client := &Client{
		client:      httpClient,
		accessToken: accessToken,
		baseURL:     defaultBaseURL,
		apiKey:      apiKey,
		onBehalfOf:  onBehalfOf,
	}

	//Services
	client.OAuth = &oauthService{}
	client.Candidates = &candidateService{client: client}
	client.CurrentUser = &currentUserService{client: client}
	client.Jobs = &jobService{client: client}
	client.TrackingLinks = &trackingLinkService{client: client}
	return client
}

// ReadJSON reads the json value into the v param. Can only read once!
func readJSON(r io.ReadCloser, v interface{}) error {
	decoder := json.NewDecoder(r)
	err := decoder.Decode(&v)
	return err
}

// Params are used to send parameters with the request.
type Params map[string]interface{}

// newRequest creates an authenticated API request that is ready to send.
func (c *Client) newRequest(method string, endpoint string, params Params, body interface{}) (*http.Request, error) {
	method = strings.ToUpper(method)
	requestURL := fmt.Sprintf("%sv1/%s", c.baseURL, endpoint)

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
		err := json.NewEncoder(&buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, requestURL, &buf)

	if c.accessToken != "" {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
	}
	if c.apiKey != "" && c.onBehalfOf != "" {
		enc := base64.StdEncoding.EncodeToString([]byte(c.apiKey))
		req.Header.Set("Authorization", fmt.Sprintf("Basic %s:", enc))
		req.Header.Set("On-Behalf-Of", c.onBehalfOf)

	}

	if req.Method == "POST" || req.Method == "PUT" {
		req.Header.Set("Content-Type", "application/json; charset=utf-8")
	}

	return req, err
}

// do takes a prepared API request and makes the API call to Greenhouse.
// It will decode the JSON into a destination struct you provide as well
// as parse any validation errors that may have occurred.
// It returns a Response object that provides a wrapper around http.Response
// with some convenience methods.
func (c *Client) do(req *http.Request, v interface{}) error {
	return do(c.client, req, v)
}

func do(client *http.Client, req *http.Request, v interface{}) error {
	req.Close = true
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	if resp == nil {
		return ErrShouldNotBeNil
	}
	defer resp.Body.Close()

	if r, err := isError(resp); r && err == nil {
		if r, err = isClientError(resp); r && err == nil {
			clientError := ClientError{
				StatusCode: resp.StatusCode,
			}
			err = readJSON(resp.Body, &clientError)
			if err != nil {
				return err
			}
			return clientError
		}
		if r, err = isServerError(resp); r && err == nil {
			serverError := ServerError{
				StatusCode: resp.StatusCode,
			}
			err = readJSON(resp.Body, &serverError)
			if err != nil {
				return err
			}
			return serverError
		}
	} else if err != nil {
		return err
	}

	err = readJSON(resp.Body, &v)
	return err
}

func interfaceToCSV(a interface{}) string {
	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(a)), ","), "[]")
}
func spaceDelimit(a interface{}) string {
	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(a)), " "), "[]")
}
