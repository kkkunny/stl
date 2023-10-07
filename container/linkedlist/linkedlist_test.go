package linkedlist

import (
	"testing"

	stltest "github.com/kkkunny/stl/test"
)

func TestNewLinkedList(t *testing.T) {
	v := NewLinkedList[int]()
	stltest.AssertEq(t, v.Length(), 0)
}

func TestNewLinkedListWith(t *testing.T) {
	v := NewLinkedListWith(1, 2, 3)
	stltest.AssertEq(t, v.Length(), 3)
}

func TestLinkedLis_Push(t *testing.T) {
	v := NewLinkedList[int]()
	v.PushBack(1)
	stltest.AssertEq(t, v.Back(), 1)
	v.PushFront(2)
	stltest.AssertEq(t, v.Front(), 2)
}

func TestLinkedList_Pop(t *testing.T) {
	v := NewLinkedListWith(1)
	v.PushFront(0)
	stltest.AssertEq(t, v.PopFront(), 0)
	stltest.AssertEq(t, v.PopBack(), 1)
}

func TestLinkedList_Clear(t *testing.T) {
	v := NewLinkedListWith(1, 2, 3)
	v.Clear()
	stltest.AssertEq(t, int(v.Length()), 0)
}

func TestLinkedList_Empty(t *testing.T) {
	v := NewLinkedListWith(1, 2, 3)
	v.Clear()
	stltest.AssertEq(t, v.Empty(), true)
}
