package stlerror

import (
	"errors"
	"fmt"
	"runtime"
)

// Error 异常
type Error interface {
	error
	fmt.Stringer
	Frame() runtime.Frame
	Stacks() []runtime.Frame
	Unwrap() error
}

// 封装
func wrap(err error) Error {
	if err == nil {
		return nil
	}
	var e Error
	if errors.As(err, &e) {
		return e
	}
	return _NewError(1, err)
}

// ErrorWrap 封装异常
func ErrorWrap(err error) Error {
	return wrap(err)
}

// ErrorWith 封装异常
func ErrorWith[T any](v T, err error) (T, Error) {
	return v, wrap(err)
}

// ErrorWith2 封装异常
func ErrorWith2[T, E any](v1 T, v2 E, err error) (T, E, Error) {
	return v1, v2, wrap(err)
}

// ErrorWith3 封装异常
func ErrorWith3[T, E, F any](v1 T, v2 E, v3 F, err error) (T, E, F, Error) {
	return v1, v2, v3, wrap(err)
}

// Errorf 新建异常
func Errorf(f string, a ...any) Error {
	return _NewError(1, fmt.Errorf(f, a...))
}
