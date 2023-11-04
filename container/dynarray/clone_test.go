package dynarray

import (
	"testing"

	stltest "github.com/kkkunny/stl/test"
)

func TestDynArray_Clone(t *testing.T) {
	v1 := NewDynArrayWith[int](1, 2, 3)
	v2 := v1.Clone()
	stltest.AssertEq(t, v1.Equal(v2), true)
}
