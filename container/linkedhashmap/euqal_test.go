package linkedhashmap

import (
	"testing"

	stltest "github.com/kkkunny/stl/test"
)

func TestLinkedHashMap_Equal(t *testing.T) {
	hm1 := NewLinkedHashMapWith[int, int](1, 1, 2, 2)
	hm2 := NewLinkedHashMapWith[int, int](2, 2, 1, 1)
	hm3 := NewLinkedHashMapWith[int, int](1, 2, 2, 1)
	stltest.AssertNotEq(t, hm1, hm2)
	stltest.AssertNotEq(t, hm1, hm3)
	stltest.AssertNotEq(t, hm2, hm3)
}
