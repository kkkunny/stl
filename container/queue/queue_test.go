package queue

import (
	"fmt"
	"testing"

	"github.com/kkkunny/stl/clone"
	stltest "github.com/kkkunny/stl/test"
)

func Test_Queue_String(t *testing.T) {
	v1 := _NewQueueWith[int](1, 2, 3)
	stltest.AssertEq(t, v1.String(), "Queue{1, 2, 3}")
}

func Test_Queue_PushAndPop(t *testing.T) {
	v := _NewQueueWith[int](1)
	v.Push(2, 3)
	stltest.AssertEq(t, v.Pop(), 1)
	stltest.AssertEq(t, v.Pop(), 2)
	stltest.AssertEq(t, v.Pop(), 3)
}

func Test_Queue_Peek(t *testing.T) {
	v := _NewQueueWith[int](1, 2, 3)
	stltest.AssertEq(t, v.Peek(), 1)
	stltest.AssertEq(t, v.Pop(), 1)
	stltest.AssertEq(t, v.Peek(), 2)
}

func Test_Queue_Clear(t *testing.T) {
	v := _NewQueueWith[int](1, 2, 3)
	v.Clear()
	stltest.AssertEq(t, v.Length(), 0)
}

func Test_Queue_Empty(t *testing.T) {
	v := _NewQueueWith[int](1, 2, 3)
	stltest.AssertEq(t, v.Empty(), false)
	v.Clear()
	stltest.AssertEq(t, v.Empty(), true)
}

func Test_Queue_ToSlice(t *testing.T) {
	v := _NewQueueWith[int](1, 2, 3)
	stltest.AssertEq(t, fmt.Sprintf("%v", v.ToSlice()), fmt.Sprintf("%v", []int{1, 2, 3}))
}

func Test_Queue_Length(t *testing.T) {
	v := _NewQueueWith[int](1, 2, 3)
	stltest.AssertEq(t, v.Length(), 3)
}

func Test_Queue_Equal(t *testing.T) {
	v1 := _NewQueueWith[int](1, 2, 3)
	v2 := _NewQueueWith[int](1, 2, 3)
	v3 := _NewQueueWith[int](3, 2, 1)
	stltest.AssertEq(t, v1.Equal(v2), true)
	stltest.AssertEq(t, v2.Equal(v3), false)
	stltest.AssertEq(t, v1.Equal(v3), false)
}

func Test_Queue_Clone(t *testing.T) {
	v1 := _NewQueueWith[int](1, 2, 3)
	v2 := clone.Clone(v1)
	stltest.AssertEq(t, v1.Equal(v2), true)
}
