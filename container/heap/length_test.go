package heap

import (
	"testing"

	stltest "github.com/kkkunny/stl/test"
)

func TestHeap_Length(t *testing.T) {
	h := NewMinHeapWith(3, 2, 1)
	stltest.AssertEq(t, h.Length(), 3)
}
