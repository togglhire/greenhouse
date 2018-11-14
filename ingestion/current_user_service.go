package ingestion

var _ CurrentUserService = &currentUserService{}

type CurrentUserService interface {
	Retrieve() (User, error)
}

type currentUserService struct {
	client *Client
}

func (cu *currentUserService) Retrieve() (result User, err error) {
	req, err := cu.client.newRequest("GET", "partner/current_user", nil, nil)
	if err != nil {
		return
	}
	err = cu.client.do(req, &result)
	return
}
