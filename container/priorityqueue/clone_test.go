package queue

import (
	"testing"

	stltest "github.com/kkkunny/stl/test"
)

func TestPriorityQueue_Clone(t *testing.T) {
	v1 := NewPriorityQueue[int]()
	v1.Push(1, 1)
	v2 := v1.Clone()
	stltest.AssertEq(t, v1.Equal(v2), true)
}
