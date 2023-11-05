package treemap

import (
	"testing"

	stltest "github.com/kkkunny/stl/test"
)

func TestTreeMap_Length(t *testing.T) {
	hm := NewTreeMapWith[int, int](1, 1, 2, 2)
	stltest.AssertEq(t, hm.Length(), 2)
}
