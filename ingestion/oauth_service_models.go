package ingestion

type OAuthScope string

const (
	OAuthScopeCandidatesCreate = OAuthScope("candidates.create")
	OAuthScopeCandidatesView   = OAuthScope("candidates.view")
	OAuthScopeJobsView         = OAuthScope("jobs.view")
)
