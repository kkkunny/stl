package stack

import (
	"fmt"
	"testing"

	stltest "github.com/kkkunny/stl/test"
)

func Test_Stack_String(t *testing.T) {
	v := _NewStackWith[int](1, 2, 3)
	stltest.AssertEq(t, v.String(), "Stack{1, 2, 3}")
}

func Test_Stack_PushAndPop(t *testing.T) {
	v := _NewStackWith[int](1)
	v.Push(2, 3)
	stltest.AssertEq(t, v.Pop(), 3)
	stltest.AssertEq(t, v.Pop(), 2)
	stltest.AssertEq(t, v.Pop(), 1)
}

func Test_Stack_Peek(t *testing.T) {
	v := _NewStackWith[int](1, 2, 3)
	stltest.AssertEq(t, v.Peek(), 3)
	stltest.AssertEq(t, v.Pop(), 3)
	stltest.AssertEq(t, v.Peek(), 2)
}

func Test_Stack_Clear(t *testing.T) {
	v := _NewStackWith[int](1, 2, 3)
	v.Clear()
	stltest.AssertEq(t, v.Length(), 0)
}

func Test_Stack_Empty(t *testing.T) {
	v := _NewStackWith[int](1, 2, 3)
	stltest.AssertEq(t, v.Empty(), false)
	v.Clear()
	stltest.AssertEq(t, v.Empty(), true)
}

func Test_Stack_ToSlice(t *testing.T) {
	v := _NewStackWith[int](1, 2, 3)
	stltest.AssertEq(t, fmt.Sprintf("%v", v.ToSlice()), fmt.Sprintf("%v", []int{1, 2, 3}))
}

func Test_Stack_Length(t *testing.T) {
	v := _NewStackWith[int](1, 2, 3)
	stltest.AssertEq(t, v.Length(), 3)
}

func Test_Stack_Equal(t *testing.T) {
	v1 := _NewStackWith[int](1, 2, 3)
	v2 := _NewStackWith[int](1, 2, 3)
	v3 := _NewStackWith[int](3, 2, 1)
	stltest.AssertEq(t, v1.Equal(v2), true)
	stltest.AssertEq(t, v2.Equal(v3), false)
	stltest.AssertEq(t, v1.Equal(v3), false)
}

func Test_Stack_Clone(t *testing.T) {
	v1 := _NewStackWith[int](1, 2, 3)
	v2 := v1.Clone()
	stltest.AssertEq(t, v1.Equal(v2), true)
}
