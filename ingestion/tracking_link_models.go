package ingestion

type PostTrackingLinkResponse struct {
	TrackingLink string `json:"tracking_link"`
	Job          string `json:"job"`
	Source       string `json:"source"`
	Referrer     string `json:"referrer"`
}
