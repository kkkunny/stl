package linkedlist

import (
	"testing"

	"github.com/kkkunny/stl/container/iterator"
	stltest "github.com/kkkunny/stl/test"
)

func TestLinkedList_Iterator(t *testing.T) {
	v1 := NewLinkedListWith[int](1, 2, 3)
	v2 := iterator.Map[int, int, LinkedList[int]](v1, func(v int) int {
		return v - 1
	})
	var i int
	for iter := v2.Iterator(); iter.Next(); {
		stltest.AssertEq(t, iter.Value(), i)
		i++
	}
}