package linkedhashmap

import (
	"testing"

	stltest "github.com/kkkunny/stl/test"
)

func TestLinkedHashMap_Capacity(t *testing.T) {
	hm := NewLinkedHashMapWithCapacity[int, int](20)
	stltest.AssertEq(t, hm.Capacity(), 20)
}
