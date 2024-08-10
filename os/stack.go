package stlos

import (
	"runtime"

	stlslices "github.com/kkkunny/stl/container/slices"
)

// GetCallStacks 获取调用栈，第一个为当前栈，最后一个为最浅的栈
func GetCallStacks(depth uint, skip ...uint) []runtime.Frame {
	skipVal := stlslices.Last(skip, 0)

	var stacks []runtime.Frame
	depth += 2
	pcs := make([]uintptr, depth)
	n := runtime.Callers(int(skipVal)+2, pcs)
	frames := runtime.CallersFrames(pcs[:n-1])
	for frame, exist := frames.Next(); exist; frame, exist = frames.Next() {
		stacks = append(stacks, frame)
	}
	return stacks
}

// GetCurrentCallStack 获取当前栈
func GetCurrentCallStack(skip ...uint) runtime.Frame {
	return stlslices.Last(GetCallStacks(1, stlslices.First(skip, 0)+1))
}
