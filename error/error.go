package stlerror

import (
	"runtime"

	stlos "github.com/kkkunny/stl/os"
)

type _Error struct {
	stacks []runtime.Frame
	err    error
}

func _NewError(skip uint, err error) *_Error {
	return &_Error{
		stacks: stlos.GetCallStacks(100, skip+1),
		err:    err,
	}
}

func (self _Error) Error() string {
	return self.err.Error()
}

func (self _Error) String() string {
	return self.err.Error()
}

// Stacks 获取栈帧信息
func (self _Error) Stacks() []runtime.Frame {
	return self.stacks
}

// Frame 获取栈帧信息
func (self _Error) Frame() runtime.Frame {
	return self.stacks[len(self.stacks)-1]
}

// Unwrap 解除封装
func (self _Error) Unwrap() error {
	return self.err
}
