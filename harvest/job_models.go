package harvest

type JobStatus string

const (
	JobStatusOpen   JobStatus = "open"
	JobStatusClosed JobStatus = "closed"
	JobStatusDraft  JobStatus = "draft"
)

type Job struct {
	Id     int64  `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

type JobListParams struct {
	PerPage              int64  `url:"per_page,omitempty"`
	Page                 int64  `url:"page,omitempty"`
	CreatedBefore        string `url:"created_before,omitempty"`
	CreatedAfter         string `url:"created_after,omitempty"`
	UpdatedBefore        string `url:"updated_before,omitempty"`
	UpdatedAfter         string `url:"updated_after,omitempty"`
	RequisitionId        string `url:"requisition_id,omitempty"`
	OpeningId            string `url:"opening_id,omitempty"`
	Status               string `url:"status,omitempty"`
	DepartmentId         string `url:"department_id,omitempty"`
	ExternalDepartmentId string `url:"external_department_id,omitempty"`
	OfficeId             string `url:"office_id,omitempty"`
	ExternalOfficeId     string `url:"external_office_id,omitempty"`
	CustomFieldOptionId  string `url:"custom_field_option_id,omitempty"`
}
