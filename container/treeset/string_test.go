package treeset

import (
	"testing"

	stltest "github.com/kkkunny/stl/test"
)

func TestTreeSet_String(t *testing.T) {
	hm := NewTreeSetWith[int](2, 1)
	stltest.AssertEq(t, hm.String(), "TreeSet{1, 2}")
}
