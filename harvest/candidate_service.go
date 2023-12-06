package harvest

import (
	"fmt"
	"net/http"
)

const (
	CANDIDATES  = "candidates"
	ATTACHMENTS = "attachments"
	NOTES       = "activity_feed/notes"
)

type CandidatesService interface {
	List(CandidateListParams) ([]Candidate, error)
	Retrieve(int64) (*Candidate, error)
	Add(*Candidate) error
	Alter(int64, *Candidate) error
	AddAttachment(int, *Attachment) error
	AddNote(int, *Note) error
}

type candidateService struct {
	client *Client
}

func NewCandidatesService(client *Client) *candidateService {
	return &candidateService{client}
}

// Slice of Candidates or Slice of pointers to Candidates?
func (s *candidateService) List(queryParams CandidateListParams) ([]Candidate, error) {
	params := structToURLValues(queryParams)
	request, err := s.client.newRequest(http.MethodGet, CANDIDATES, params, nil)
	if err != nil {
		return nil, err
	}

	var candidates []Candidate
	err = s.client.do(request, candidates)

	return candidates, err
}

func (s *candidateService) Retrieve(id int64) (*Candidate, error) {
	candidate := &Candidate{}
	request, err := s.client.newRequest(http.MethodGet, fmt.Sprintf("%s/%d", CANDIDATES, id), nil, nil)
	if err != nil {
		return nil, err
	}

	err = s.client.do(request, candidate)

	return candidate, err
}

func (s *candidateService) Add(candidate *Candidate) error {
	request, err := s.client.newRequest(http.MethodPost, CANDIDATES, nil, candidate)
	if err != nil {
		return err
	}

	return s.client.do(request, candidate)
}

// NOTE: Id could also be in the candidate struct
func (s *candidateService) Alter(id int64, candidate *Candidate) error {
	request, err := s.client.newRequest(http.MethodPut, fmt.Sprintf("%s/%d", CANDIDATES, id), nil, candidate)
	if err != nil {
		return err
	}

	return s.client.do(request, candidate)
}

func (s *candidateService) AddAttachment(id int, attachment *Attachment) error {
	request, err := s.client.newRequest(http.MethodPost, fmt.Sprintf("%s/%d/%s", CANDIDATES, id, ATTACHMENTS), nil, attachment)
	if err != nil {
		return err
	}

	return s.client.do(request, attachment)
}

func (s *candidateService) AddNote(id int, note *Note) error {
	request, err := s.client.newRequest(http.MethodPost, fmt.Sprintf("%s/%d/%s", CANDIDATES, id, NOTES), nil, note)
	if err != nil {
		return err
	}

	return s.client.do(request, note)
}
