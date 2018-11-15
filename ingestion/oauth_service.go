package ingestion

import (
	"net/url"
)

var _ OAuthService = &oauthService{}

type OAuthService interface {
	CreateAuthURL(consumerKey string, scopes []OAuthScope, state string) (string, error)
}

type oauthService struct {
}

func NewOAuthService() OAuthService {
	return &oauthService{}
}

func (o *oauthService) CreateAuthURL(consumerKey string, scopes []OAuthScope, state string) (result string, err error) {
	if consumerKey == "" {
		return "", ErrConsumerKeyMissing
	}
	greenhouseAuthURL, err := url.Parse("https://api.greenhouse.io/oauth/authorize")
	if err != nil {
		return "", err
	}
	q := greenhouseAuthURL.Query()
	q.Add("client_id", consumerKey)
	if len(scopes) != 0 {
		q.Add("scope", spaceDelimit(scopes))
	}
	if state != "" {
		q.Add("state", state)
	}
	greenhouseAuthURL.RawQuery = q.Encode()
	return greenhouseAuthURL.String(), nil
}
