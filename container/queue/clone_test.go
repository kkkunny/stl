package queue

import (
	"testing"

	stltest "github.com/kkkunny/stl/test"
)

func TestDynArray_Clone(t *testing.T) {
	v1 := NewQueueWith[int](1, 2, 3)
	v2 := v1.Clone()
	stltest.AssertEq(t, v1.Equal(v2), true)
}
