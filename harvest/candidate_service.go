package harvest

import (
	"fmt"
	"net/http"

	"github.com/google/go-querystring/query"
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
	Edit(int64, *Candidate) error
	AddAttachment(int64, *Attachment) error
	AddNote(int64, *Note) error
}

type candidateService struct {
	client *Client
}

func NewCandidatesService(client *Client) *candidateService {
	return &candidateService{client}
}

// Slice of Candidates or Slice of pointers to Candidates?
func (s *candidateService) List(queryParams CandidateListParams) ([]Candidate, error) {
	params, err := query.Values(queryParams)
	if err != nil {
		return nil, NewSDKError(fmt.Sprintf("error parsing query params: %s", err.Error()))
	}
	request, err := s.client.newRequest(http.MethodGet, CANDIDATES, params, nil)
	if err != nil {
		return nil, err
	}

	var candidates []Candidate
	err = s.client.do(request, &candidates)

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
func (s *candidateService) Edit(id int64, candidate *Candidate) error {
	request, err := s.client.newRequest(http.MethodPatch, fmt.Sprintf("%s/%d", CANDIDATES, id), nil, candidate)
	if err != nil {
		return err
	}

	return s.client.do(request, candidate)
}

func (s *candidateService) AddAttachment(id int64, attachment *Attachment) error {
	request, err := s.client.newRequest(http.MethodPost, fmt.Sprintf("%s/%d/%s", CANDIDATES, id, ATTACHMENTS), nil, attachment)
	if err != nil {
		return err
	}

	return s.client.do(request, attachment)
}

func (s *candidateService) AddNote(id int64, note *Note) error {
	request, err := s.client.newRequest(http.MethodPost, fmt.Sprintf("%s/%d/%s", CANDIDATES, id, NOTES), nil, note)
	if err != nil {
		return err
	}

	return s.client.do(request, note)
}
