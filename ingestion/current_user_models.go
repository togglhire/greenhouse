package ingestion

// User is the struct used for retreiving the current user from Greenhouse
type User struct {
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Email     string `json:"email,omitempty"`
}
