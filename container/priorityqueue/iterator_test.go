package queue

import (
	"testing"

	"github.com/kkkunny/stl/container/iterator"
	"github.com/kkkunny/stl/container/pair"
	stltest "github.com/kkkunny/stl/test"
)

func TestPriorityQueue_Iterator(t *testing.T) {
	v1 := NewPriorityQueue[int]()
	v1.Push(1, 2)
	v2 := iterator.Map[pair.Pair[uint64, int], pair.Pair[uint64, int], PriorityQueue[int]](v1, func(v pair.Pair[uint64, int]) pair.Pair[uint64, int] {
		return pair.NewPair(v.First, v.Second-1)
	})
	var i int
	for iter := v2.Iterator(); iter.Next(); {
		stltest.AssertEq(t, iter.Value(), pair.NewPair[uint64, int](1, 1))
		i++
	}
}
