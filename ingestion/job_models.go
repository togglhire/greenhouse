package ingestion

type Job struct {
	ID     int64  `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	Status string `json:"status,omitempty"`
	Public bool   `json:"public,omitempty"`
}
