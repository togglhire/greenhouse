package ingestion

var _ TrackingLinkService = &trackingLinkService{}

type TrackingLinkService interface {
	Post(jobID int64) (PostTrackingLinkResponse, error)
}

type trackingLinkService struct {
	client *Client
}

func (tl *trackingLinkService) Post(jobID int64) (result PostTrackingLinkResponse, err error) {
	c := struct {
		JobID int64 `json:"job_id"`
	}{
		JobID: jobID,
	}
	req, err := tl.client.newRequest("POST", "partner/tracking_link", nil, c)
	if err != nil {
		return
	}
	err = tl.client.do(req, &result)
	return
}
