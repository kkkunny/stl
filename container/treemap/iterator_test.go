package treemap

import (
	"testing"

	"github.com/kkkunny/stl/container/iterator"
	"github.com/kkkunny/stl/container/pair"
	stltest "github.com/kkkunny/stl/test"
)

func TestTreeMap_Iterator(t *testing.T) {
	v1 := NewTreeMapWith[int, int](1, 1, 2, 2)
	v2 := iterator.Map[pair.Pair[int, int], pair.Pair[int, int], TreeMap[int, int]](v1, func(v pair.Pair[int, int]) pair.Pair[int, int] {
		return pair.NewPair(v.First-1, v.Second-1)
	})
	var i int
	for iter := v2.Iterator(); iter.Next(); {
		stltest.AssertEq(t, iter.Value().First, i)
		i++
	}
}
