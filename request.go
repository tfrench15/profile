package profile

// Request represents a request that will be made to the Profile API.
type Request interface {
	// Validate ensures the request is valid.
	Validate() error

	// internal ensures no other package can create a type that is a valid request.
	internal()
}
