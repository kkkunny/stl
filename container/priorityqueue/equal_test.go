package queue

import (
	"testing"

	stltest "github.com/kkkunny/stl/test"
)

func TestPriorityQueue_Equal(t *testing.T) {
	v1 := NewPriorityQueue[int]()
	v1.Push(1, 1)
	v2 := NewPriorityQueue[int]()
	v2.Push(1, 1)
	v3 := NewPriorityQueue[int]()
	v3.Push(2, 2)
	stltest.AssertEq(t, v1.Equal(v2), true)
	stltest.AssertEq(t, v2.Equal(v3), false)
	stltest.AssertEq(t, v1.Equal(v3), false)
}
