package treemap

import (
	"testing"

	stltest "github.com/kkkunny/stl/test"
)

func TestTreeMap_String(t *testing.T) {
	hm := NewTreeMapWith[int, int](2, 2, 1, 1)
	stltest.AssertEq(t, hm.String(), "TreeMap{1: 1, 2: 2}")
}
