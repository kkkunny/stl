package btree

import (
	"testing"

	stltest "github.com/kkkunny/stl/test"
)

func TestNewBTree(t *testing.T) {
	v := NewBTree[int]()
	stltest.AssertEq(t, v.Length(), 0)
}

func TestBTree_Push(t *testing.T) {
	v := NewBTree[int]()
	stltest.AssertEq(t, v.Push(1).Value, 1)
	stltest.AssertEq(t, v.Push(2).Value, 2)
}

func TestBTree_Remove(t *testing.T) {
	v := NewBTree[int]()
	v.Push(2)
	v.Push(1)
	v.Push(3)
	stltest.AssertEq(t, v.Remove(2).Value, 2)
	stltest.AssertEq(t, v.Top().Value, 3)
}
