package queue

import (
	"testing"

	stltest "github.com/kkkunny/stl/test"
)

func TestQueue_Length(t *testing.T) {
	v := NewQueueWith[int](1, 2, 3)
	stltest.AssertEq(t, v.Length(), 3)
}
