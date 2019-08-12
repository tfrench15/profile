package profile

// PersonasError is returned when the Profile API returns an error.
type PersonasError struct {
	Err *ErrorCode
}

// ErrorCode details the reason the error occurred.
type ErrorCode struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// Error implements the error interface for PersonasError.
func (e *PersonasError) Error() string {
	return "CODE:\t" + e.Err.Code + "\nMESSAGE:\t" + e.Err.Message
}
