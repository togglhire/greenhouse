package harvest

import "fmt"

const (
	JOBS = "jobs"
)

type JobsService interface {
	// An alternative would be allowing users to pass a map of string to interface{}.
	List(ListJobsQueryParams) ([]Job, error)
	Retrieve(int64) (*Job, error)
}

type jobService struct {
	client *Client
}

func NewJobsService(client *Client) *jobService {
	return &jobService{client}
}

func (s *jobService) List(params ListJobsQueryParams) ([]Job, error) {
	request, err := s.client.newRequest("GET", JOBS, structToURLValues(params), nil)
	if err != nil {
		return nil, err
	}

	jobs := make([]Job, 0)
	if err = s.client.do(request, jobs); err != nil {
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
