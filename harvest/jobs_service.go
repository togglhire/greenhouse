package harvest

type JobsService interface {
	// List() ([]Job, error)
	Retrieve(int64) (*Job, error)
}

type jobService struct {
	client *Client
}

func NewJobsService(client *Client) *jobService {
	return &jobService{client}
}

func (s *jobService) Retrieve(Id int64) (*Job, error) {
	_, err := s.client.newRequest("GET", "jobs/"+string(Id), nil, nil)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
