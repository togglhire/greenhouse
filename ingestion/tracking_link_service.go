package ingestion

var _ TrackingLinkService = &trackingLinkService{}

type TrackingLinkService interface {
	Post(jobID int64) ([]TrackingLinkResponse, error)
}

type trackingLinkService struct {
	client *Client
}

func (tl *trackingLinkService) Post(jobID int64) (result []TrackingLinkResponse, err error) {
	c := struct {
		JobID int64 `json:"job_id"`
	}{
		JobID: jobID,
	}
	req, err := tl.client.newRequest("POST", "partner/candidates", nil, c)
	if err != nil {
		return
	}
	err = tl.client.do(req, &result)
	return
}
