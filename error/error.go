package stlerror

import (
	"runtime"
)

type _Error struct {
	stacks []runtime.Frame
	err    error
}

func _NewError(skip uint, err error) *_Error {
	var reverseStacks []runtime.Frame
	pcs := make([]uintptr, 20)

	n := runtime.Callers(int(skip)+2, pcs)
	frames := runtime.CallersFrames(pcs[:n-1])
	for frame, exist := frames.Next(); exist; frame, exist = frames.Next() {
		if !exist {
			break
		}
		reverseStacks = append(reverseStacks, frame)
	}

	stacks := make([]runtime.Frame, len(reverseStacks))
	for i, s := range reverseStacks {
		stacks[len(reverseStacks)-i-1] = s
	}

	return &_Error{
		stacks: stacks,
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
