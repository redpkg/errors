package errors

// NewStatusError .
func NewStatusError(statusCode, code int, message string) *StatusError {
	return &StatusError{
		StatusCode: statusCode,
		Code:       code,
		Message:    message,
	}
}

// StatusError .
type StatusError struct {
	StatusCode int    `json:"-"`
	Code       int    `json:"code"`
	Message    string `json:"message"`
	Internal   error  `json:"-"`
}

// Error .
func (e *StatusError) Error() string {
	return e.Message
}

// Unwrap .
func (e *StatusError) Unwrap() error {
	return e.Internal
}

// SetInternal .
func (e *StatusError) SetInternal(err error) *StatusError {
	e.Internal = err
	return e
}
