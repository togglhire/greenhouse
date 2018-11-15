package ingestion

var _ CandidateService = &candidateService{}

type CandidateService interface {
	Retrieve(ids []int64) ([]Candidate, error)
	Post(candidates []PostCandidate) ([]PostCandidateResponse, error)
}

type candidateService struct {
	client *Client
}

func (s *candidateService) Retrieve(ids []int64) (result []Candidate, err error) {
	params := Params{
		"candidate_ids": interfaceToCSV(ids),
	}
	req, err := s.client.newRequest("GET", "partner/candidates", params, nil)
	if err != nil {
		return
	}
	err = s.client.do(req, &result)
	return
}

func (s *candidateService) Post(candidates []PostCandidate) (result []PostCandidateResponse, err error) {
	req, err := s.client.newRequest("POST", "partner/candidates", nil, candidates)
	if err != nil {
		return
	}
	err = s.client.do(req, &result)
	return
}
