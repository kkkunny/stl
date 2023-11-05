package linkedhashmap

import (
	"testing"

	stltest "github.com/kkkunny/stl/test"
)

func TestHashMap_Length(t *testing.T) {
	hm := NewLinkedHashMapWith[int, int](1, 1, 2, 2)
	stltest.AssertEq(t, hm.Length(), 2)
}
