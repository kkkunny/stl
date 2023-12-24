package queue

import (
	"testing"

	stltest "github.com/kkkunny/stl/test"
)

func TestPriorityQueue_Length(t *testing.T) {
	v := NewPriorityQueue[int]()
	stltest.AssertEq(t, v.Length(), 0)
	v.Push(1, 1)
	stltest.AssertEq(t, v.Length(), 1)
	v.Pop()
	stltest.AssertEq(t, v.Length(), 0)
}
