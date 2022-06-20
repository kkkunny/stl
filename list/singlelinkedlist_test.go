package list

import (
	"github.com/kkkunny/stl/util"
	"testing"
)

func TestSingleLinkedList_PopBack(t *testing.T) {
	l := NewSingleLinkedList[int](1, 2, 3)
	util.AssertEq(t, l.PopBack(), 3)
	util.AssertEq(t, l.PopBack(), 2)
	util.AssertEq(t, l.PopBack(), 1)
}
