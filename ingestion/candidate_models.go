package ingestion

// Candidate is the struct used for retreiving candidates from Greenhouse
type Candidate struct {
	ID           int64         `json:"id,omitempty"`
	Name         string        `json:"name,omitempty"`
	ExternalID   string        `json:"external_id,omitempty"`
	Applications []Application `json:"applications,omitempty"`
}

// PostCandidate is the struct used for creating a new candidate in Greenhouse
type PostCandidate struct {
	Prospect            bool          `json:"prospect"`
	FirstName           string        `json:"first_name,omitempty"`
	LastName            string        `json:"last_name,omitempty"`
	Company             string        `json:"company,omitempty"`
	Title               string        `json:"title,omitempty"`
	Resume              string        `json:"resume,omitempty"`
	JobID               int64         `json:"job_id,omitempty"`
	ExternalID          string        `json:"external_id,omitempty"`
	Notes               string        `json:"notes,omitempty"`
	ProspectPoolID      int64         `json:"prospect_pool_id,omitempty"`
	ProspectPoolStageID int64         `json:"prospect_pool_stage_id,omitempty"`
	ProspectOwnerEmail  string        `json:"prospect_owner_email,omitempty"`
	PhoneNumbers        []PhoneNumber `json:"phone_numbers,omitempty"`
	Emails              []Email       `json:"emails,omitempty"`
	SocialMedia         []SocialMedia `json:"social_media,omitempty"`
	Websites            []Website     `json:"websites,omitempty"`
	Addresses           []Address     `json:"addresses,omitempty"`
}

// PostCandidateResponse is the struct that is returned when a new candidate is created in Greenhouse
type PostCandidateResponse struct {
	ID            int64  `json:"id,omitempty"`
	ApplicationID int64  `json:"application_id,omitempty"`
	ExternalID    string `json:"external_id,omitempty"`
	ProfileURL    string `json:"profile_url,omitempty"`
}

// PhoneNumber struct is for representing a phone number in greenhouse
type PhoneNumber struct {
	PhoneNumber string          `json:"phone_number,omitempty"`
	Type        PhoneNumberType `json:"type,omitempty"`
}

// PhoneNumberType is for representing the possible types of phone numbers
type PhoneNumberType string

const (
	PhoneNumberTypeMobile = PhoneNumberType("mobile")
	PhoneNumberTypeHome   = PhoneNumberType("home")
	PhoneNumberTypeWork   = PhoneNumberType("work")
	PhoneNumberTypeOther  = PhoneNumberType("other")
)

// Email struct is for representing an email address in greenhouse
type Email struct {
	Email string    `json:"email,omitempty"`
	Type  EmailType `json:"type,omitempty"`
}

type EmailType string

const (
	EmailTypePersonal = EmailType("personal")
	EmailTypeWork     = EmailType("work")
	EmailTypeOther    = EmailType("other")
)

type Address struct {
	Address string      `json:"address,omitempty"`
	Type    AddressType `json:"type,omitempty"`
}

type AddressType string

const (
	AddressTypeHome  = AddressType("home")
	AddressTypeWork  = AddressType("work")
	AddressTypeOther = AddressType("other")
)

type SocialMedia struct {
	URL string `json:"url,omitempty"`
}

type Website struct {
	URL  string      `json:"url,omitempty"`
	Type WebsiteType `json:"type,omitempty"`
}

type WebsiteType string

const (
	WebsiteTypePersonal  = WebsiteType("personal")
	WebsiteTypeCompany   = WebsiteType("company")
	WebsiteTypePortfolio = WebsiteType("portfolio")
	WebsiteTypeBlog      = WebsiteType("blog")
	WebsiteTypeOther     = WebsiteType("other")
)
