package stlerr

import (
	"github.com/pkg/errors"
)

type StackTracer interface {
	StackTrace() errors.StackTrace
}

type errorWithStack interface {
	error
	StackTrace() errors.StackTrace
	Unwrap() error
}

type _Error struct {
	err   error
	stack errors.StackTrace
}

func (e _Error) Error() string {
	return e.err.Error()
}

func (e _Error) String() string {
	return e.err.Error()
}

func (e _Error) StackTrace() errors.StackTrace {
	return e.stack
}

func (e _Error) Unwrap() error {
	return e.err
}
