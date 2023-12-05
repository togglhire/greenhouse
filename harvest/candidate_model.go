package harvest

import (
	"net/url"
	"reflect"
)

type Candidate struct {
	Id                   int64        `json:"id,omitempty"`
	FirstName            string       `json:"first_name"`
	LastName             string       `json:"last_name"`
	Company              string       `json:"company,omitempty"`
	Title                string       `json:"title,omitempty"`
	PhoneNumbers         []KeyValue   `json:"phone_numbers,omitempty"`
	Addresses            []KeyValue   `json:"address,omitempty"`
	EmailAddresses       []KeyValue   `json:"email_addresses,omitempty"`
	WebsiteAddresses     []KeyValue   `json:"website_addresses,omitempty"`
	SocialMediaAddresses []KeyValue   `json:"social_media_addresses,omitempty"`
	Educations           []Education  `json:"educations,omitempty"`
	Employments          []Employment `json:"employments,omitempty"`
	Tags                 []string     `json:"tags,omitempty"`
	// still missing some fields
}

// Can't be KeyValue because of the type field
type KeyValue struct {
	Value string `json:"value"`
	Type  string `json:"type,omitempty"`
}

type Education struct {
	Id           int64  `json:"school_id,omitempty"`
	DisciplineId int64  `json:"discipline_id,omitempty"`
	DegreeId     int64  `json:"degree_id,omitempty"`
	StartDate    string `json:"start_date,omitempty"`
	EndDate      string `json:"end_date,omitempty"`
}

type Employment struct {
	CompanyName int64  `json:"company_name,omitempty"`
	Title       string `json:"title,omitempty"`
	StartDate   string `json:"start_date,omitempty"`
	EndDate     string `json:"end_date,omitempty"`
}

type ListCandidatesQueryParams struct {
	PerPage       int64  `url:"per_page,omitempty"`
	Page          int64  `url:"page,omitempty"`
	CreatedBefore string `url:"created_before,omitempty"`
	CreatedAfter  string `url:"created_after,omitempty"`
	UpdatedBefore string `url:"updated_before,omitempty"`
	UpdatedAfter  string `url:"updated_after,omitempty"`
	JobId         int64  `url:"job_id,omitempty"`
	Email         string `url:"email,omitempty"`
	CandidateIds  string `url:"candidate_ids,omitempty"`
}

// utils
func ListCandidatesQueryParamsToURLValues(data interface{}) url.Values {
	values := url.Values{}
	val := reflect.ValueOf(data)
	typ := reflect.TypeOf(data)

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		tag := typ.Field(i).Tag.Get("url")
		if field.IsZero() && tag == "omitempty" {
			continue
		}

		values.Add(tag, field.String())
	}

	return values
}
