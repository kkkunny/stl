package linkedlist

import (
	"testing"

	stltest "github.com/kkkunny/stl/test"
)

func TestLinkedList_String(t *testing.T) {
	v := NewLinkedListWith[int](1, 2, 3)
	stltest.AssertEq(t, v.String(), "LinkedList{1, 2, 3}")
}
