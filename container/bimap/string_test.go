package bimap

import (
	"testing"

	stltest "github.com/kkkunny/stl/test"
)

func TestBiMap_String(t *testing.T) {
	hm := NewBiMapWith[int, int](1, 1, 2, 2)
	stltest.AssertNotEq(t, hm.String(), "")
}
