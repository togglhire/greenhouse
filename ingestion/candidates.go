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
	params := Params{
		"candidate_ids": int64ArrayToCSV,
	}
	req, err := s.client.newRequest("GET", "partner/candidates", params, nil)
	if err != nil {
		return
	}
	err = s.client.do(req, &candidates)
	return
}

func (s *candidateService) Post(candidates []PostCandidate) ([]PostCandidateResponse, error) {
	return []PostCandidateResponse{}, ErrNotImplemented
}
