package errors_test

import (
	"net/http"
	"testing"

	"github.com/redpkg/errors/v2"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	assert := assert.New(t)

	err := errors.New(50000, "server error").
		SetInternal(errors.New(40000, "foobar"))

	assert.Equal(http.StatusInternalServerError, err.StatusCode)
	assert.Equal(50000, err.Code)
	assert.Equal("server error", err.Message)
	assert.EqualError(err, "server error")
	assert.EqualError(err.Unwrap(), "foobar")

	err.SetStatusCode(http.StatusNotImplemented)
	assert.Equal(http.StatusNotImplemented, err.StatusCode)
}

func TestFlatten(t *testing.T) {
	assert := assert.New(t)

	err1 := errors.New(1, "error 1")
	err2 := errors.New(2, "error 2").SetInternal(err1)
	err3 := errors.New(3, "error 3").SetInternal(err2)

	assert.EqualValues([]error{err3, err2, err1}, errors.Flatten(err3))
}
