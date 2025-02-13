package errorapp

import (
	"fmt"
)

type StatusError struct {
	error
	status int
}

func (se StatusError) StatusCode() int {
	return se.status
}

func WithStatus(status int, msg string, args ...any) error {
	return StatusError{
		error:  fmt.Errorf(msg, args...),
		status: status,
	}
}
