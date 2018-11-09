package ingestion

type Application struct {
	ID         int64  `json:"id"`
	Job        string `json:"job"`
	Status     string `json:"status"`
	Stage      string `json:"stage"`
	ProfileURL string `json:"profile_url"`
}
