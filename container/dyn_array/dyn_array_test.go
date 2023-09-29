package dynarray

import (
	"testing"

	stltest "github.com/kkkunny/stl/test"
)

func TestNewDynArray(t *testing.T) {
	v := NewDynArray[int]()
	stltest.AssertEq(t, v.Length(), 0)
	stltest.AssertEq(t, v.Capacity(), 0)
}

func TestNewDynArrayWithCapacity(t *testing.T) {
	v := NewDynArrayWithCapacity[int](10)
	stltest.AssertEq(t, v.Capacity(), 10)
}

func TestNewDynArrayWithLength(t *testing.T) {
	v := NewDynArrayWithLength[int](10)
	stltest.AssertEq(t, v.Length(), 10)
	stltest.AssertEq(t, v.Capacity(), 10)
}

func TestNewDynArrayWith(t *testing.T) {
	v := NewDynArrayWith(1, 2, 3)
	stltest.AssertEq(t, v.Length(), 3)
	stltest.AssertEq(t, v.Capacity(), 3)
}

func TestDynArray_Get(t *testing.T) {
	v := NewDynArrayWith(1, 2, 3)
	stltest.AssertEq(t, v.Get(0), 1)
	stltest.AssertEq(t, v.Get(1), 2)
	stltest.AssertEq(t, v.Get(2), 3)
}

func TestDynArray_Set(t *testing.T) {
	v := NewDynArrayWithLength[int](1)
	stltest.AssertEq(t, v.Set(0, 1), 0)
	stltest.AssertEq(t, v.Get(0), 1)
}

func TestDynArray_PushBack(t *testing.T) {
	v := NewDynArray[int]()
	v.PushBack(1)
	stltest.AssertEq(t, v.Get(0), 1)
}

func TestDynArray_PushFront(t *testing.T) {
	v := NewDynArrayWith(1)
	v.PushFront(0)
	stltest.AssertEq(t, v.Get(0), 0)
	stltest.AssertEq(t, v.Get(1), 1)
}

func TestDynArray_Insert(t *testing.T) {
	v := NewDynArrayWith(1, 3)
	v.Insert(1, 2)
	stltest.AssertEq(t, v.Get(0), 1)
	stltest.AssertEq(t, v.Get(1), 2)
	stltest.AssertEq(t, v.Get(2), 3)
}

func TestDynArray_Remove(t *testing.T) {
	v := NewDynArrayWith(1, 2, 3)
	stltest.AssertEq(t, v.Remove(1), 2)
	stltest.AssertEq(t, v.Get(0), 1)
	stltest.AssertEq(t, v.Get(1), 3)
}

func TestDynArray_PopBack(t *testing.T) {
	v := NewDynArrayWith(1, 2, 3)
	stltest.AssertEq(t, v.PopBack(), 3)
	stltest.AssertEq(t, v.PopBack(), 2)
	stltest.AssertEq(t, v.PopBack(), 1)
}

func TestDynArray_PopFront(t *testing.T) {
	v := NewDynArrayWith(1, 2, 3)
	stltest.AssertEq(t, v.PopFront(), 1)
	stltest.AssertEq(t, v.PopFront(), 2)
	stltest.AssertEq(t, v.PopFront(), 3)
}

func TestDynArray_Back(t *testing.T) {
	v := NewDynArrayWith(1, 2, 3)
	stltest.AssertEq(t, v.Back(), 3)
}

func TestDynArray_Front(t *testing.T) {
	v := NewDynArrayWith(1, 2, 3)
	stltest.AssertEq(t, v.Front(), 1)
}

func TestDynArray_Clear(t *testing.T) {
	v := NewDynArrayWith(1, 2, 3)
	v.Clear()
	stltest.AssertEq(t, int(v.Length()), 0)
}

func TestDynArray_Empty(t *testing.T) {
	v := NewDynArrayWith(1, 2, 3)
	v.Clear()
	stltest.AssertEq(t, v.Empty(), true)
}
