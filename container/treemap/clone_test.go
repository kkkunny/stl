package treemap

import (
	"testing"

	stltest "github.com/kkkunny/stl/test"
)

func TestTreeMap_Clone(t *testing.T) {
	hm1 := NewTreeMapWith[int, int](1, 1, 2, 2)
	hm2 := hm1.Clone()
	stltest.AssertEq(t, hm1, hm2)
}
