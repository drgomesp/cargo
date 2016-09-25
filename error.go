package cargo

import (
	"errors"
	"fmt"
)

// Error represents an error with context information
type Error struct {
	Err error
}

// NewError creates a new contextual error
func NewError(msg string, args ...interface{}) Error {
	return Error{
		Err: errors.New(fmt.Errorf(msg, args...).Error()),
	}
}

func (e Error) Error() string {
	return e.Err.Error()
}
