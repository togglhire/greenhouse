package ingestion

import "net/url"

type OAuthScope string

const (
	OAuthScopeCandidatesCreate = OAuthScope("candidates.create")
	OAuthScopeCandidatesView   = OAuthScope("candidates.view")
	OAuthScopeJobsView         = OAuthScope("jobs.view")
)

// AuthURLData holds the info required to create an AuthURL.
type AuthURLData struct {
	ConsumerKey string
	Scopes      []OAuthScope
	RedirectURI string
	State       string
}

// AccessTokenData holds the info required to retrieve the access token.
type AccessTokenData struct {
	ConsumerKey    string
	ConsumerSecret string
	RedirectURI    string
	RequestURI     *url.URL
}
