package bimap

import (
	"testing"

	stltest "github.com/kkkunny/stl/test"
)

func TestBiMap_Equal(t *testing.T) {
	hm1 := NewBiMapWith[int, int](1, 1, 2, 2)
	hm2 := NewBiMapWith[int, int](2, 2, 1, 1)
	hm3 := NewBiMapWith[int, int](1, 2, 2, 1)
	stltest.AssertEq(t, hm1, hm2)
	stltest.AssertNotEq(t, hm1, hm3)
	stltest.AssertNotEq(t, hm2, hm3)
}
