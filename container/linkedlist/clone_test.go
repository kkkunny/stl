package linkedlist

import (
	"testing"

	stltest "github.com/kkkunny/stl/test"
)

func TestLinkedList_Clone(t *testing.T) {
	v1 := NewLinkedListWith[int](1, 2, 3)
	v2 := v1.Clone()
	stltest.AssertEq(t, v1, v2)
}
