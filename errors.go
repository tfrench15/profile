package profile

// PersonasError is returned when the Profile API returns an error.
type PersonasError struct {
	Error *ErrorCode
}

// ErrorCode details the reason the error occurred.
type ErrorCode struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
