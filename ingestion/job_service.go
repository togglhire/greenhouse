package ingestion

var _ JobService = &jobService{}

type JobService interface {
	Retrieve() ([]Job, error)
}

type jobService struct {
	client *Client
}

func (j *jobService) Retrieve() (result []Job, err error) {
	req, err := j.client.newRequest("GET", "partner/jobs", nil, nil)
	if err != nil {
		return
	}
	err = j.client.do(req, &result)
	return
}
