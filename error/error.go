package stlerror

import (
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

type _Error struct {
	stacks []runtime.Frame
	err    error
}

func (e _Error) Error() string {
	return e.err.Error()
}

func (e _Error) String() string {
	return e.err.Error()
}

// Stacks 获取栈帧信息
func (e _Error) Stacks() []runtime.Frame {
	return e.stacks
}

// Frame 获取栈帧信息
func (e _Error) Frame() runtime.Frame {
	return e.stacks[0]
}

// Unwrap 解除封装
func (e _Error) Unwrap() error {
	return e.err
}
