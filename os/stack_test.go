package stlos

import (
	"strings"
	"testing"

	stlslices "github.com/kkkunny/stl/container/slices"
	stltest "github.com/kkkunny/stl/test"
)

func TestGetCallStacks(t *testing.T) {
	stacks := GetCallStacks(1)
	stltest.AssertEq(t, stlslices.Last(strings.Split(stlslices.First(stacks).Function, ".")), "TestGetCallStacks")
}

func TestGetCurrentCallStack(t *testing.T) {
	stack := GetCurrentCallStack()
	stltest.AssertEq(t, stlslices.Last(strings.Split(stack.Function, ".")), "TestGetCurrentCallStack")
}
