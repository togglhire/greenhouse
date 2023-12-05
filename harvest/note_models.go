package harvest

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
