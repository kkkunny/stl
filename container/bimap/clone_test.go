package bimap

import (
	"testing"

	stltest "github.com/kkkunny/stl/test"
)

func TestBiMap_Clone(t *testing.T) {
	hm1 := NewBiMapWith[int, int](1, 1, 2, 2)
	hm2 := hm1.Clone()
	stltest.AssertEq(t, hm1, hm2)
}
