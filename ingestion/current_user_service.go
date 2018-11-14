package ingestion

var _ CurrentUserService = &currentUserService{}

type CurrentUserService interface {
	Retrieve() (User, error)
}

type currentUserService struct {
	client *Client
}

func (cu *currentUserService) Retrieve() (result User, err error) {
	return
}
