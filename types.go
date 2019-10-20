package profile

// Cursor holds pagination information concerning the requested data.
type Cursor struct {
	URL     string `json:"url"`
	HasMore bool   `json:"has_more"`
	Next    string `json:"next"`
}

// Request represents a request that will be made to the Profile API.
type Request interface {
	// Validate ensures the request is valid.
	Validate() error

	// internal ensures no other package can create a type that is a valid request.
	internal()
}

// Response represents possible responses from calling Query() on a Request.
type Response interface {
	// Marshal marshals the underlying data type into JSON.
	Marshal() ([]byte, error)

	// internal ensures no other package can create a type that is a valid request.
	internal()
}
