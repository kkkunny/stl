package stlerr

import (
	stlos "github.com/kkkunny/stl/os"
)

type StackError interface {
	StackFrames() []stlos.Frame
}

type _Error struct {
	err   error
	stack []stlos.Frame
}

func WithStack(err error, frames []stlos.Frame) error {
	return &_Error{
		err:   err,
		stack: frames,
	}
}

func (e _Error) Error() string {
	return e.err.Error()
}

func (e _Error) String() string {
	return e.err.Error()
}

func (e _Error) StackFrames() []stlos.Frame {
	return e.stack
}

func (e _Error) Unwrap() error {
	return e.err
}
