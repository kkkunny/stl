package linkedlist

import (
	"fmt"
	"testing"

	stltest "github.com/kkkunny/stl/test"
)

func TestLinkedList_PushBack(t *testing.T) {
	v := NewLinkedListWith[int](1)
	v.PushBack(2)
	stltest.AssertEq(t, v.Front(), 1)
	stltest.AssertEq(t, v.Back(), 2)
}

func TestLinkedList_PushFront(t *testing.T) {
	v := NewLinkedListWith[int](1)
	v.PushBack(2)
	stltest.AssertEq(t, v.Front(), 1)
	stltest.AssertEq(t, v.Back(), 2)
}

func TestLinkedList_PopBack(t *testing.T) {
	v := NewLinkedListWith[int](1, 2, 3)
	stltest.AssertEq(t, v.PopBack(), 3)
	stltest.AssertEq(t, v.Front(), 1)
	stltest.AssertEq(t, v.Back(), 2)
}

func TestLinkedList_PopFront(t *testing.T) {
	v := NewLinkedListWith[int](0, 1, 2)
	stltest.AssertEq(t, v.PopFront(), 0)
	stltest.AssertEq(t, v.Front(), 1)
	stltest.AssertEq(t, v.Back(), 2)
}

func TestLinkedList_Back(t *testing.T) {
	v := NewLinkedListWith[int](1, 2, 3)
	stltest.AssertEq(t, v.Back(), 3)
}

func TestLinkedList_Front(t *testing.T) {
	v := NewLinkedListWith[int](1, 2, 3)
	stltest.AssertEq(t, v.Front(), 1)
}

func TestLinkedList_Clear(t *testing.T) {
	v := NewLinkedListWith[int](1, 2, 3)
	v.Clear()
	stltest.AssertEq(t, v.Length(), 0)
}

func TestLinkedList_Empty(t *testing.T) {
	v := NewLinkedListWith[int](1, 2, 3)
	stltest.AssertEq(t, v.Empty(), false)
	v.Clear()
	stltest.AssertEq(t, v.Empty(), true)
}

func TestLinkedList_Append(t *testing.T) {
	var v LinkedList[int]
	v.Append(NewLinkedListWith(1, 2))
	stltest.AssertEq(t, v.Front(), 1)
	stltest.AssertEq(t, v.Back(), 2)
}

func TestLinkedList_ToSlice(t *testing.T) {
	v := NewLinkedListWith(1, 2, 3)
	stltest.AssertEq(t, fmt.Sprintf("%v", v.ToSlice()), fmt.Sprintf("%v", []int{1, 2, 3}))
}
