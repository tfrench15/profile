package profile

// Cursor holds pagination information concerning the requested data.
type Cursor struct {
	URL     string `json:"url"`
	HasMore bool   `json:"has_more"`
	Next    string `json:"next"`
}

// Metadata holds the metadata for the requested Profile.
type Metadata struct {
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
	ExpiresAt      string `json:"expires_at"`
	FirstMessageID string `json:"first_message_id"`
	FirstSourceID  string `json:"first_source_id"`
	LastMessageID  string `json:"last_message_id"`
	LastSourceID   string `json:"last_source_id"`
}

// ExternalID represents an external ID in the Personas Identity Graph.
type ExternalID struct {
	SourceID       string `json:"source_id"`
	Collection     string `json:"collection"`
	ID             string `json:"id"`
	Type           string `json:"type"`
	CreatedAt      string `json:"created_at"`
	Encoding       string `json:"encoding"`
	FirstMessageID string `json:"first_message_id"`
}
