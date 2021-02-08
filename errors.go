package errors

import "net/http"

// New .
func New(code int, message string) *Error {
	return &Error{
		StatusCode: http.StatusInternalServerError,
		Code:       code,
		Message:    message,
	}
}

// Error .
type Error struct {
	StatusCode int    `json:"-"`
	Code       int    `json:"code"`
	Message    string `json:"message"`
	Internal   error  `json:"-"`
}

// Error .
func (e *Error) Error() string {
	return e.Message
}

// Unwrap .
func (e *Error) Unwrap() error {
	return e.Internal
}

// SetStatusCode .
func (e *Error) SetStatusCode(code int) *Error {
	e.StatusCode = code
	return e
}

// SetInternal .
func (e *Error) SetInternal(err error) *Error {
	e.Internal = err
	return e
}

// Flatten .
func Flatten(err error) []error {
	errs := []error{}

	for {
		errs = append(errs, err)

		x, ok := err.(interface{ Unwrap() error })
		if !ok {
			return errs
		}

		if err = x.Unwrap(); err == nil {
			return errs
		}
	}
}
