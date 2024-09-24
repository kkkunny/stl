package stack

import (
	"testing"

	stltest "github.com/kkkunny/stl/test"
)

func TestStack_String(t *testing.T) {
	v1 := NewStackWith[int](1, 2, 3)
	stltest.AssertEq(t, v1.String(), "StackFrames{1, 2, 3}")
}
