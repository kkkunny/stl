package stlerr

import (
	"github.com/pkg/errors"

	stlslices "github.com/kkkunny/stl/container/slices"
	stlos "github.com/kkkunny/stl/os"
)

func GetErrorStackFrames(err error) []stlos.Frame {
	if err == nil {
		return nil
	}

	type StackTracer interface {
		StackTrace() errors.StackTrace
	}
	if e, ok := err.(StackTracer); ok {
		return stlslices.Map(e.StackTrace(), func(_ int, f errors.Frame) stlos.Frame {
			return stlos.WrapErrorFrame(f)
		})
	}

	if e, ok := err.(StackError); ok {
		return e.StackFrames()
	}

	if e, ok := err.(interface{ Unwrap() error }); ok {
		return GetErrorStackFrames(e.Unwrap())
	}

	return nil
}
