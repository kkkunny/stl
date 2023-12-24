package queue

import (
	"testing"

	stltest "github.com/kkkunny/stl/test"
)

func TestPriorityQueue_String(t *testing.T) {
	v := NewPriorityQueue[int]()
	v.Push(1, 1)
	v.Push(2, 2)
	stltest.AssertEq(t, v.String(), "PriorityQueue{2:2, 1:1}")
}
