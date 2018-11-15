package ingestion

import (
	"net/http"
	"net/url"
)

var _ OAuthService = &oauthService{}

type OAuthService interface {
	CreateAuthURL(AuthURLData) (string, error)
	GetAccessToken(AccessTokenData) (accessToken string, err error)
}

type oauthService struct {
	client         *http.Client
	consumerKey    string
	consumerSecret string
}

func NewOAuthService(consumerKey, consumerSecret string, httpClient *http.Client) OAuthService {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	return &oauthService{
		client:         httpClient,
		consumerKey:    consumerKey,
		consumerSecret: consumerSecret,
	}
}

func (o *oauthService) CreateAuthURL(d AuthURLData) (result string, err error) {
	if d.ConsumerKey == "" && o.consumerKey == "" {
		return "", ErrConsumerKeyMissing
	}
	if d.ConsumerKey == "" {
		d.ConsumerKey = o.consumerKey
	}
	greenhouseAuthURL, err := url.Parse(authorizeURL)
	if err != nil {
		return "", err
	}
	q := greenhouseAuthURL.Query()
	q.Add("client_id", d.ConsumerKey)
	if len(d.Scopes) != 0 {
		q.Add("scope", spaceDelimit(d.Scopes))
	}
	if d.RedirectURI != "" {
		q.Add("redirect_uri", d.RedirectURI)
	}
	if d.State != "" {
		q.Add("state", d.State)
	}
	q.Add("response_type", "code")
	greenhouseAuthURL.RawQuery = q.Encode()
	return greenhouseAuthURL.String(), nil
}

func (o *oauthService) GetAccessToken(d AccessTokenData) (accessToken string, err error) {
	if d.RequestURI == nil {
		return "", ErrRequestURINil
	}
	code := d.RequestURI.Query().Get("code")
	if len(code) == 0 {
		return "", ErrCodeMissing
	}
	if d.ConsumerKey == "" {
		d.ConsumerKey = o.consumerKey
	}
	if d.ConsumerSecret == "" {
		d.ConsumerSecret = o.consumerSecret
	}

	greenhouseAccessTokenURL, err := url.Parse(accessTokenURL)
	if err != nil {
		return "", err
	}
	q := greenhouseAccessTokenURL.Query()
	q.Add("grant_type", "authorization_code")
	q.Add("code", code)
	q.Add("client_id", d.ConsumerKey)
	q.Add("client_secret", d.ConsumerSecret)
	q.Add("redirect_uri", d.RedirectURI)
	greenhouseAccessTokenURL.RawQuery = q.Encode()

	result := struct {
		AccessToken string `json:"access_token,omitempty"`
	}{}
	req, err := http.NewRequest("POST", greenhouseAccessTokenURL.String(), nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("Accept", "application/json")

	err = do(o.client, req, &result)
	return result.AccessToken, err
}
