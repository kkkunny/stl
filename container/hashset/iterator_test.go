package hashset

import (
	"testing"

	"github.com/kkkunny/stl/container/iterator"
	stltest "github.com/kkkunny/stl/test"
)

func TestHashSet_Iterator(t *testing.T) {
	v1 := NewHashSetWith[int](1, 2, 2)
	v2 := iterator.Map[int, int, HashSet[int]](v1, func(v int) int {
		return v - 1
	})
	stltest.AssertEq(t, v2, NewHashSetWith[int](0, 1))
}
