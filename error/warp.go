package stlerr

import (
	"fmt"
	"runtime"

	stlslices "github.com/kkkunny/stl/container/slices"
	stlos "github.com/kkkunny/stl/os"
)

// 封装
func wrap(err error) error {
	if err == nil {
		return nil
	}

	stacks := GetErrorStackFrames(err)
	if len(stacks) == 0 {
		stacks = stlslices.Map(stlos.GetCallStacks(32, 2), func(_ int, f runtime.Frame) stlos.Frame {
			return stlos.WrapRuntimeFrame(f)
		})
	}

	return WithStack(err, stacks)
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
	return WithStack(fmt.Errorf(f, a...), stlslices.Map(stlos.GetCallStacks(32, 1), func(_ int, f runtime.Frame) stlos.Frame {
		return stlos.WrapRuntimeFrame(f)
	}))
}
