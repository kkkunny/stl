package queue

import (
	"testing"

	stltest "github.com/kkkunny/stl/test"
)

func TestDynArray_String(t *testing.T) {
	v1 := NewQueueWith[int](1, 2, 3)
	stltest.AssertEq(t, v1.String(), "Queue{1, 2, 3}")
}
