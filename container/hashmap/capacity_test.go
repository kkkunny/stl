package hashmap

import (
	"testing"

	stltest "github.com/kkkunny/stl/test"
)

func TestHashMap_Capacity(t *testing.T) {
	var hm HashMap[int, int]
	stltest.AssertEq(t, hm.Capacity(), initialBucketCapacity)
	hm = NewHashMapWithCapacity[int, int](20)
	stltest.AssertEq(t, hm.Capacity(), 20)
}
