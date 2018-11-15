package ingestion

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_oauthService_CreateAuthURL(t *testing.T) {
	type args struct {
		consumerKey string
		scopes      []OAuthScope
		state       string
	}
	test := struct {
		args    args
		wantURL string
		wantErr bool
	}{
		args: args{
			consumerKey: "consumer-key",
			scopes:      []OAuthScope{OAuthScopeCandidatesCreate, OAuthScopeCandidatesView, OAuthScopeJobsView},
		},
		wantURL: "https://api.greenhouse.io/oauth/authorize?client_id=consumer-key&scope=candidates.create+candidates.view+jobs.view",
	}

	gotURL, err := NewOAuthService().CreateAuthURL(test.args.consumerKey, test.args.scopes, test.args.state)

	switch test.wantErr {
	case true:
		assert.Error(t, err)
	case false:
		assert.NoError(t, err)
	}
	assert.Equal(t, test.wantURL, gotURL)
}
