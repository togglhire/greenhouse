package ingestion

import (
	"log"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_oauthService_CreateAuthURL(t *testing.T) {
	type args struct {
		data AuthURLData
	}
	test := struct {
		args    args
		wantURL string
		wantErr bool
	}{
		args: args{
			data: AuthURLData{
				ConsumerKey: "consumer-key",
				Scopes:      []OAuthScope{OAuthScopeCandidatesCreate, OAuthScopeCandidatesView, OAuthScopeJobsView},
				RedirectURI: "https://example.com",
				State:       "some-state",
			},
		},
		wantURL: "https://api.greenhouse.io/oauth/authorize?client_id=consumer-key&redirect_uri=https%3A%2F%2Fexample.com&response_type=code&scope=candidates.create+candidates.view+jobs.view&state=some-state",
	}

	gotURL, err := NewOAuthService("", "", nil).CreateAuthURL(test.args.data)

	switch test.wantErr {
	case true:
		assert.Error(t, err)
	case false:
		assert.NoError(t, err)
	}
	assert.Equal(t, test.wantURL, gotURL)
}

func Test_oauthService_GetAccessToken(t *testing.T) {
	t.Skip("Requires working consumer key, secret & token to work")
	requestURI, err := url.Parse("https://example.com/api/v1/integrations/greenhouse/callback?code=Mbg0SEV2L2&state=somestate")
	assert.NoError(t, err)
	type args struct {
		data AccessTokenData
	}
	test := struct {
		args    args
		wantErr bool
	}{
		args: args{
			data: AccessTokenData{
				ConsumerKey:    "consumer-key",
				ConsumerSecret: "consumer-secret",
				RequestURI:     requestURI,
				RedirectURI:    "https://example.com/api/v1/integrations/greenhouse/callback",
			},
		}}

	accessToken, err := NewOAuthService("", "", nil).GetAccessToken(test.args.data)

	switch test.wantErr {
	case true:
		assert.Error(t, err)
	case false:
		assert.NoError(t, err)
	}
	log.Printf("accessToken: %#+v\n", accessToken)
}
