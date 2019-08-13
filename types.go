package profile

// Cursor holds pagination information concerning the requested data.
type Cursor struct {
	URL     string `json:"url"`
	HasMore bool   `json:"has_more"`
	Next    string `json:"next"`
}
