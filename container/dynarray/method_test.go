package dynarray

import (
	"fmt"
	"testing"

	stltest "github.com/kkkunny/stl/test"
)

func TestDynArray_Get(t *testing.T) {
	v := NewDynArrayWith[int](1, 2, 3)
	stltest.AssertEq(t, v.Get(0), 1)
	stltest.AssertEq(t, v.Get(1), 2)
	stltest.AssertEq(t, v.Get(2), 3)
}

func TestDynArray_Set(t *testing.T) {
	v := NewDynArrayWith[int](1, 2, 3)
	stltest.AssertEq(t, v.Set(0, 3), 1)
	stltest.AssertEq(t, v.Get(0), 3)
}

func TestDynArray_PushBack(t *testing.T) {
	v := NewDynArrayWith[int](1)
	v.PushBack(2, 3)
	stltest.AssertEq(t, v.Get(0), 1)
	stltest.AssertEq(t, v.Get(1), 2)
	stltest.AssertEq(t, v.Get(2), 3)
}

func TestDynArray_PushFront(t *testing.T) {
	v := NewDynArrayWith[int](3)
	v.PushFront(1, 2)
	stltest.AssertEq(t, v.Get(0), 1)
	stltest.AssertEq(t, v.Get(1), 2)
	stltest.AssertEq(t, v.Get(2), 3)
}

func TestDynArray_Insert(t *testing.T) {
	v := NewDynArrayWith[int](1, 4)
	v.Insert(1, 2, 3)
	stltest.AssertEq(t, v.Get(0), 1)
	stltest.AssertEq(t, v.Get(1), 2)
	stltest.AssertEq(t, v.Get(2), 3)
	stltest.AssertEq(t, v.Get(3), 4)
}

func TestDynArray_Remove(t *testing.T) {
	v := NewDynArrayWith[int](1, 2, 4, 3)
	stltest.AssertEq(t, v.Remove(2), 4)
	stltest.AssertEq(t, v.Get(0), 1)
	stltest.AssertEq(t, v.Get(1), 2)
	stltest.AssertEq(t, v.Get(2), 3)
}

func TestDynArray_PopBack(t *testing.T) {
	v := NewDynArrayWith[int](1, 2, 3, 4)
	stltest.AssertEq(t, v.PopBack(), 4)
	stltest.AssertEq(t, v.Get(0), 1)
	stltest.AssertEq(t, v.Get(1), 2)
	stltest.AssertEq(t, v.Get(2), 3)
}

func TestDynArray_PopFront(t *testing.T) {
	v := NewDynArrayWith[int](0, 1, 2, 3)
	stltest.AssertEq(t, v.PopFront(), 0)
	stltest.AssertEq(t, v.Get(0), 1)
	stltest.AssertEq(t, v.Get(1), 2)
	stltest.AssertEq(t, v.Get(2), 3)
}

func TestDynArray_Back(t *testing.T) {
	v := NewDynArrayWith[int](1, 2, 3)
	stltest.AssertEq(t, v.Back(), 3)
}

func TestDynArray_Front(t *testing.T) {
	v := NewDynArrayWith[int](1, 2, 3)
	stltest.AssertEq(t, v.Front(), 1)
}

func TestDynArray_Clear(t *testing.T) {
	v := NewDynArrayWith[int](1, 2, 3)
	v.Clear()
	stltest.AssertEq(t, v.Length(), 0)
}

func TestDynArray_Empty(t *testing.T) {
	v := NewDynArrayWith[int](1, 2, 3)
	stltest.AssertEq(t, v.Empty(), false)
	v.Clear()
	stltest.AssertEq(t, v.Empty(), true)
}

func TestDynArray_Append(t *testing.T) {
	var v DynArray[int]
	v.Append(NewDynArrayWith(1, 2, 3))
	stltest.AssertEq(t, v.Get(0), 1)
	stltest.AssertEq(t, v.Get(1), 2)
	stltest.AssertEq(t, v.Get(2), 3)
}

func TestDynArray_ToSlice(t *testing.T) {
	v := NewDynArrayWith(1, 2, 3)
	stltest.AssertEq(t, fmt.Sprintf("%v", v.ToSlice()), fmt.Sprintf("%v", []int{1, 2, 3}))
}
