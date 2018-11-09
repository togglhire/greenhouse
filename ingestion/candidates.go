package ingestion

var _ CandidateService = &candidateService{}

type Candidate struct {
	ID           int64         `json:"id"`
	Name         string        `json:"name"`
	ExternalID   string        `json:"external_id"`
	Applications []Application `json:"applications"`
}

type CandidateService interface {
	Get(ids []int64) (candidates []Candidate, err error)
}

type candidateService struct {
	client *Client
}

func (s *candidateService) Get(ids []int64) (candidates []Candidate, err error) {
	return
}
