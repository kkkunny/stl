package stack

import (
	"testing"

	stltest "github.com/kkkunny/stl/test"
)

func TestDynArray_Equal(t *testing.T) {
	v1 := NewStackWith[int](1, 2, 3)
	v2 := NewStackWith[int](1, 2, 3)
	v3 := NewStackWith[int](3, 2, 1)
	stltest.AssertEq(t, v1.Equal(v2), true)
	stltest.AssertEq(t, v2.Equal(v3), false)
	stltest.AssertEq(t, v1.Equal(v3), false)
}
