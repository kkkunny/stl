package stlerr

import (
	"fmt"

	"github.com/pkg/errors"

	stlos "github.com/kkkunny/stl/os"
)

// 封装
func wrap(err error) errorWithStack {
	if err == nil {
		return nil
	}
	if e, ok := errors.Cause(err).(StackTracer); ok {
		return &_Error{
			stack: e.StackTrace(),
			err:   err,
		}
	}
	return &_Error{
		err:   err,
		stack: stlos.GetErrorTrace(2),
	}
}

// ErrorWrap 封装异常
func ErrorWrap(err error) error {
	return wrap(err)
}

// ErrorWith 封装异常
func ErrorWith[T any](v T, err error) (T, error) {
	return v, wrap(err)
}

// ErrorWith2 封装异常
func ErrorWith2[T, E any](v1 T, v2 E, err error) (T, E, error) {
	return v1, v2, wrap(err)
}

// ErrorWith3 封装异常
func ErrorWith3[T, E, F any](v1 T, v2 E, v3 F, err error) (T, E, F, error) {
	return v1, v2, v3, wrap(err)
}

// Errorf 新建异常
func Errorf(f string, a ...any) error {
	return &_Error{
		err:   fmt.Errorf(f, a...),
		stack: stlos.GetErrorTrace(1),
	}
}
