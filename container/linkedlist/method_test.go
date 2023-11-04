package linkedlist

import (
	"testing"

	stltest "github.com/kkkunny/stl/test"
)

func TestDynArray_PushBack(t *testing.T) {
	v := NewLinkedListWith[int](1)
	v.PushBack(2)
	stltest.AssertEq(t, v.Front(), 1)
	stltest.AssertEq(t, v.Back(), 2)
}

func TestDynArray_PushFront(t *testing.T) {
	v := NewLinkedListWith[int](1)
	v.PushBack(2)
	stltest.AssertEq(t, v.Front(), 1)
	stltest.AssertEq(t, v.Back(), 2)
}

func TestDynArray_PopBack(t *testing.T) {
	v := NewLinkedListWith[int](1, 2, 3)
	stltest.AssertEq(t, v.PopBack(), 3)
	stltest.AssertEq(t, v.Front(), 1)
	stltest.AssertEq(t, v.Back(), 2)
}

func TestDynArray_PopFront(t *testing.T) {
	v := NewLinkedListWith[int](0, 1, 2)
	stltest.AssertEq(t, v.PopFront(), 0)
	stltest.AssertEq(t, v.Front(), 1)
	stltest.AssertEq(t, v.Back(), 2)
}

func TestDynArray_Back(t *testing.T) {
	v := NewLinkedListWith[int](1, 2, 3)
	stltest.AssertEq(t, v.Back(), 3)
}

func TestDynArray_Front(t *testing.T) {
	v := NewLinkedListWith[int](1, 2, 3)
	stltest.AssertEq(t, v.Front(), 1)
}

func TestDynArray_Clear(t *testing.T) {
	v := NewLinkedListWith[int](1, 2, 3)
	v.Clear()
	stltest.AssertEq(t, v.Length(), 0)
}

func TestDynArray_Empty(t *testing.T) {
	v := NewLinkedListWith[int](1, 2, 3)
	stltest.AssertEq(t, v.Empty(), false)
	v.Clear()
	stltest.AssertEq(t, v.Empty(), true)
}
