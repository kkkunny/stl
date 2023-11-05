package hashmap

import (
	"testing"

	stltest "github.com/kkkunny/stl/test"
)

func TestHashMap_String(t *testing.T) {
	hm := NewHashMapWith[int, int](1, 1, 2, 2)
	stltest.AssertNotEq(t, hm.String(), "")
}
