package harvest

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

const (
	DEFAULT_BASE_URL         = "https://harvest.greenhouse.io/"
	AUTH_HEADER_KEY          = "Authorization"
	CONTENT_TYPE_HEADER_KEY  = "Content-Type"
	ON_BEAHALF_OF_HEADER_KEY = "On-Behalf-Of"
)

type Version string

const (
	V1 Version = "v1"
)

// type ClientConfig struct {
// 	// HTTP client used to communicate with the Harvest API.
// 	HttpClient *http.Client

// 	// Base URL for API requests.
// 	BaseURL string

// 	// Revision of the Harvest API to use.
// 	ApiVersion Version

// 	// Harvest API apiKey
// 	ApiKey string

// 	// On-Behalf-Of header value
// 	OnBehalfOf string
// }

type Client struct {
	// HTTP client used to communicate with the Harvest API.
	client *http.Client

	// Base URL for API requests.
	baseURL string

	// Revision of the Harvest API to use.
	apiVersion Version

	// Harvest API apiKey
	apiKey string

	// On-Behalf-Of header value
	onBehalfOf string

	// Services used for talking to different parts of the Harvest API.
	Candidates CandidatesService
	Jobs       JobsService
}

func NewDefaultClient(apiKey string, onBehalfOf string, httpClient *http.Client) (*Client, error) {
	return newClient(apiKey, onBehalfOf, httpClient, DEFAULT_BASE_URL, V1)
}

func NewClient(apiKey string, onBehalfOf string, httpClient *http.Client, baseURL string, apiVersion Version) (*Client, error) {
	return newClient(apiKey, onBehalfOf, httpClient, baseURL, apiVersion)
}

func newClient(apiKey string, onBehalfOf string, httpClient *http.Client, baseURL string, apiVersion Version) (*Client, error) {
	if apiKey == "" {
		return nil, NewSDKError("api key is required")
	}

	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	client := &Client{
		client:     httpClient,
		baseURL:    baseURL + string(apiVersion) + "/",
		apiKey:     apiKey,
		onBehalfOf: onBehalfOf,
	}

	var candidatesService CandidatesService
	var jobsService JobsService
	if apiVersion == V1 {
		candidatesService = NewCandidatesService(client)
		jobsService = NewJobsService(client)
	} // else if apiVersion == V2 {
	// }

	client.Candidates = candidatesService
	client.Jobs = jobsService

	return client, nil
}

func (c *Client) newRequest(method string, endpointPath string, params url.Values, body interface{}) (*http.Request, error) {
	method = strings.ToUpper(method)
	if isValidMethod(method) {
		return nil, NewSDKError("invalid method provided")
	}
	requestURL := fmt.Sprintf("%s%s", c.baseURL, endpointPath)

	if len(params) > 0 {
		requestURL += "?" + params.Encode()
	}

	var buf bytes.Buffer
	if body != nil {
		err := json.NewEncoder(&buf).Encode(body)
		if err != nil {
			return nil, NewSDKError("error encoding body")
		}
	}

	req, err := http.NewRequest(method, requestURL, &buf)

	enc := base64.StdEncoding.EncodeToString([]byte(c.apiKey))
	req.Header.Set(AUTH_HEADER_KEY, fmt.Sprintf("Basic %s:", enc))

	if c.onBehalfOf != "" {
		req.Header.Set(ON_BEAHALF_OF_HEADER_KEY, c.onBehalfOf)
	}

	if methodSendsBody(method) {
		req.Header.Set(CONTENT_TYPE_HEADER_KEY, "application/json; charset=utf-8")
	}

	return req, err
}

func (c *Client) do(req *http.Request, v interface{}) error {
	req.Close = true
	resp, err := c.client.Do(req)
	if err != nil {
		return NewSDKError(fmt.Sprintf("error making request: %v", err))
	}
	if resp == nil {
		return NewSDKError("could not get a response, response is nil")
	}
	defer resp.Body.Close()

	isErrResp, err := isErrorResponse(resp)
	if isErrResp {
		return err
	}

	if err := readJSON(resp.Body, v); err != nil {
		return NewSDKError(fmt.Sprintf("error decoding response: %v", err))
	}

	return nil
}
