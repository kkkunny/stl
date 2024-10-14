package linkedlist

import (
	"testing"

	"github.com/kkkunny/stl/clone"
	stltest "github.com/kkkunny/stl/test"
)

func TestLinkedList_Clone(t *testing.T) {
	v1 := NewLinkedListWith[int](1, 2, 3)
	v2 := clone.Clone(v1)
	stltest.AssertEq(t, v1, v2)
}
