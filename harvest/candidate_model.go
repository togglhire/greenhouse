package harvest

type Candidate struct {
	Id                   int64                          `json:"id,omitempty"`
	FirstName            string                         `json:"first_name"`
	LastName             string                         `json:"last_name"`
	Company              string                         `json:"company,omitempty"`
	Title                string                         `json:"title,omitempty"`
	PhoneNumbers         []KeyValue[PhoneNumberType]    `json:"phone_numbers,omitempty"`
	Addresses            []KeyValue[AddressType]        `json:"address,omitempty"`
	EmailAddresses       []KeyValue[EmailAddressType]   `json:"email_addresses,omitempty"`
	WebsiteAddresses     []KeyValue[WebsiteAddressType] `json:"website_addresses,omitempty"`
	SocialMediaAddresses []KeyValue[string]             `json:"social_media_addresses,omitempty"`
	Educations           []Education                    `json:"educations,omitempty"`
	Employments          []Employment                   `json:"employments,omitempty"`
	Tags                 []string                       `json:"tags,omitempty"`
	// still missing some fields
}

type KeyValueType interface {
	PhoneNumberType | AddressType | EmailAddressType | WebsiteAddressType | string
}

type PhoneNumberType string

const (
	PNHome   PhoneNumberType = "home"
	PNWork   PhoneNumberType = "work"
	PNMobile PhoneNumberType = "mobile"
	PNSkype  PhoneNumberType = "skype"
	PNOther  PhoneNumberType = "other"
)

type AddressType string

const (
	ATHome  AddressType = "home"
	ATWork  AddressType = "work"
	ATOther AddressType = "other"
)

type EmailAddressType string

const (
	EATHome  EmailAddressType = "home"
	EATWork  EmailAddressType = "work"
	EATOther EmailAddressType = "other"
)

type WebsiteAddressType string

const (
	WATPersonal  WebsiteAddressType = "personal"
	WATCompany   WebsiteAddressType = "company"
	WATPortfolio WebsiteAddressType = "portfolio"
	WATBlog      WebsiteAddressType = "blog"
	WATOther     WebsiteAddressType = "other"
)

type KeyValue[T KeyValueType] struct {
	Value string `json:"value"`
	Type  T      `json:"type"`
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

type AttachmentType string

const (
	ATResume      AttachmentType = "resume"
	ATCoverLetter AttachmentType = "cover_letter"
	ATAdminOnly   AttachmentType = "admin_only"
)

type Attachment struct {
	Filename string         `json:"filename"`
	Type     AttachmentType `json:"type"`
	Content  string         `json:"content,omitempty"`
	URL      string         `json:"url,omitempty"`
	ContentT string         `json:"content_type,omitempty"`
}

type NoteVisibility string

const (
	NVAdminOnly NoteVisibility = "admin_only"
	NVPrivate   NoteVisibility = "private"
	NVPublic    NoteVisibility = "public"
)

type Note struct {
	UserId     int64          `json:"user,omitempty"`
	Body       string         `json:"body,omitempty"`
	Visibility NoteVisibility `json:"visibility"`
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
