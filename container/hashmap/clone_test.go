package hashmap

import (
	"testing"

	stltest "github.com/kkkunny/stl/test"
)

func TestHashMap_Clone(t *testing.T) {
	hm1 := NewHashMapWith[int, int](1, 1, 2, 2)
	hm2 := hm1.Clone()
	stltest.AssertEq(t, hm1, hm2)
}
