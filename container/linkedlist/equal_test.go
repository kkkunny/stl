package linkedlist

import (
	"testing"

	stltest "github.com/kkkunny/stl/test"
	stlval "github.com/kkkunny/stl/value"
)

func TestLinkedList_Equal(t *testing.T) {
	v1 := NewLinkedListWith[int](1, 2, 3)
	v2 := stlval.Clone(v1)
	stltest.AssertEq(t, v1, v2)
}
