package stack

import (
	"testing"

	stltest "github.com/kkkunny/stl/test"
)

func TestStack_Length(t *testing.T) {
	v := NewStackWith[int](1, 2, 3)
	stltest.AssertEq(t, v.Length(), 3)
}
