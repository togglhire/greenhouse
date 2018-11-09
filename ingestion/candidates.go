package ingestion

var _ CandidateService = &candidateService{}

type CandidateService interface {
	Retrieve(ids []int64) ([]Candidate, error)
	Post(candidates []PostCandidate) ([]PostCandidateResponse, error)
}

type candidateService struct {
	client *Client
}

func (s *candidateService) Retrieve(ids []int64) (candidates []Candidate, err error) {
	return candidates, ErrNotImplemented
}

func (s *candidateService) Post(candidates []PostCandidate) ([]PostCandidateResponse, error) {
	return []PostCandidateResponse{}, ErrNotImplemented
}
