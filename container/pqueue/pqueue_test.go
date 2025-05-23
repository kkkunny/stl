package pqueue

import (
	"fmt"
	"testing"

	"github.com/kkkunny/stl/clone"
	"github.com/kkkunny/stl/container/tuple"
	stltest "github.com/kkkunny/stl/test"
)

func TestAnyPQueue_String(t *testing.T) {
	v := _NewAnyPQueue[int, int]()
	v.Push(1, 1)
	v.Push(2, 2)
	fmt.Println(v.String())
	stltest.AssertEq(t, v.String(), "PQueue{2:2, 1:1}")
}

func TestAnyPQueue_Clone(t *testing.T) {
	v1 := _NewAnyPQueue[int, int]()
	v1.Push(1, 1)
	v2 := clone.Clone(v1)
	stltest.AssertEq(t, v1, v2)
}

func TestAnyPQueue_Equal(t *testing.T) {
	v1 := _NewAnyPQueue[int, int]()
	v1.Push(1, 1)
	v2 := _NewAnyPQueue[int, int]()
	v2.Push(1, 1)
	v3 := _NewAnyPQueue[int, int]()
	v3.Push(2, 2)
	stltest.AssertEq(t, v1.Equal(v2), true)
	stltest.AssertEq(t, v2.Equal(v3), false)
	stltest.AssertEq(t, v1.Equal(v3), false)
}

func TestAnyPQueue_Length(t *testing.T) {
	v := _NewAnyPQueue[int, int]()
	stltest.AssertEq(t, v.Length(), 0)
	v.Push(1, 1)
	stltest.AssertEq(t, v.Length(), 1)
	v.Pop()
	stltest.AssertEq(t, v.Length(), 0)
}

func TestAnyPQueue_PushAndPop(t *testing.T) {
	v := _NewAnyPQueue[int, int]()
	v.Push(1, 1)
	v.Push(2, 2)
	v.Push(3, 3)
	stltest.AssertEq(t, tuple.Pack2(v.Pop()), tuple.Pack2[int, int](3, 3))
	stltest.AssertEq(t, tuple.Pack2(v.Pop()), tuple.Pack2[int, int](2, 2))
	stltest.AssertEq(t, tuple.Pack2(v.Pop()), tuple.Pack2[int, int](1, 1))
}

func TestAnyPQueue_Peek(t *testing.T) {
	v := _NewAnyPQueue[int, int]()
	v.Push(1, 1)
	v.Push(2, 2)
	v.Push(3, 3)
	stltest.AssertEq(t, tuple.Pack2(v.Peek()), tuple.Pack2[int, int](3, 3))
	stltest.AssertEq(t, tuple.Pack2(v.Pop()), tuple.Pack2[int, int](3, 3))
	stltest.AssertEq(t, tuple.Pack2(v.Peek()), tuple.Pack2[int, int](2, 2))
}

func TestAnyPQueue_Clear(t *testing.T) {
	v := _NewAnyPQueue[int, int]()
	v.Push(1, 1)
	stltest.AssertEq(t, v.Length(), 1)
	v.Clear()
	stltest.AssertEq(t, v.Length(), 0)
}

func TestAnyPQueue_Empty(t *testing.T) {
	v := _NewAnyPQueue[int, int]()
	v.Push(1, 1)
	stltest.AssertEq(t, v.Empty(), false)
	v.Clear()
	stltest.AssertEq(t, v.Empty(), true)
}
