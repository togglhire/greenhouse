package harvest

import (
	"fmt"

	"github.com/google/go-querystring/query"
)

const (
	JOBS = "jobs"
)

type JobsService interface {
	// An alternative would be allowing users to pass a map of string to interface{}.
	List(JobListParams) ([]Job, error)
	Retrieve(int64) (*Job, error)
}

type jobService struct {
	client *Client
}

func NewJobsService(client *Client) *jobService {
	return &jobService{client}
}

func (s *jobService) List(queryParams JobListParams) ([]Job, error) {
	params, err := query.Values(queryParams)
	if err != nil {
		return nil, NewSDKError(fmt.Sprintf("error parsing query params: %s", err.Error()))
	}
	request, err := s.client.newRequest("GET", JOBS, params, nil)
	if err != nil {
		return nil, err
	}

	var jobs []Job
	if err = s.client.do(request, &jobs); err != nil {
		return nil, err
	}

	return jobs, err
}

func (s *jobService) Retrieve(id int64) (*Job, error) {
	_, err := s.client.newRequest("GET", fmt.Sprintf("%s/%d", JOBS, id), nil, nil)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
