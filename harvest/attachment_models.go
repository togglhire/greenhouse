package harvest

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
