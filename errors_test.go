package errors_test

import (
	goerrors "errors"
	"testing"

	"github.com/redpkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	assert := assert.New(t)

	err := errors.New(50000, "server error").
		SetInternal(errors.New(40000, "bad request"))

	assert.EqualError(err, "server error")
	assert.Equal(50000, err.Code)
	assert.Equal("server error", err.Message)

	assert.EqualError(err.Unwrap(), "bad request")
}

func TestNewStatusError(t *testing.T) {
	assert := assert.New(t)

	err := errors.NewStatusError(500, 50000, "server error").
		SetInternal(errors.NewStatusError(400, 40000, "bad request"))

	assert.EqualError(err, "server error")
	assert.Equal(500, err.StatusCode)
	assert.Equal(50000, err.Code)
	assert.Equal("server error", err.Message)

	assert.EqualError(err.Unwrap(), "bad request")
}

func TestFlatten(t *testing.T) {
	assert := assert.New(t)

	err1 := errors.New(1, "error 1")
	err2 := errors.NewStatusError(500, 2, "error 2")
	err3 := goerrors.New("error 3")
	err2.SetInternal(err3)
	err1.SetInternal(err2)

	assert.EqualValues([]error{err1, err2, err3}, errors.Flatten(err1))
}
