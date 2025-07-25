package stlos

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"

	"github.com/pkg/errors"

	stlslices "github.com/kkkunny/stl/container/slices"
)

// Frame 栈帧
type Frame interface {
	File() string
	Line() int
}

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

type customFrame struct {
	file string
	line int
}

func NewFrame(file string, line int) Frame {
	return &customFrame{
		file: file,
		line: line,
	}
}

func WrapErrorFrame(f errors.Frame) Frame {
	file := fmt.Sprintf("%+s", f)
	file = file[strings.LastIndex(file, "\n\t")+1:]
	line, _ := strconv.ParseInt(fmt.Sprintf("%d", f), 10, 64)
	return NewFrame(file, int(line))
}

func (f customFrame) File() string {
	return f.file
}

func (f customFrame) Line() int {
	return f.line
}

type runtimeFrame struct {
	rt runtime.Frame
}

func WrapRuntimeFrame(f runtime.Frame) Frame {
	return &runtimeFrame{rt: f}
}

func (f runtimeFrame) File() string {
	return f.rt.File
}

func (f runtimeFrame) Line() int {
	return f.rt.Line
}
