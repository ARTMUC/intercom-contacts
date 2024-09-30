package contact

type Contact struct {
	ID        string `json:"id,omitempty"`
	Email     string `json:"email,omitempty"`
	Companies any    `json:"companies,omitempty"`
}
