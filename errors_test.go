package errors_test

import (
	"testing"

	"github.com/redpkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	assert := assert.New(t)

	err := errors.New(500, "server error").
		SetInternal(errors.New(400, "bad request"))

	assert.EqualError(err, "server error")
	assert.Equal(500, err.Code)
	assert.Equal("server error", err.Message)

	assert.EqualError(err.Unwrap(), "bad request")
}

func TestFlatten(t *testing.T) {
	assert := assert.New(t)

	err1 := errors.New(1, "error 1")
	err2 := errors.New(2, "error 2")
	err3 := errors.New(3, "error 3")
	err2.SetInternal(err3)
	err1.SetInternal(err2)

	assert.EqualValues([]error{err1, err2, err3}, errors.Flatten(err1))
}
