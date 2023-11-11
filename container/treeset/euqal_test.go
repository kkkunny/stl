package treeset

import (
	"testing"

	stltest "github.com/kkkunny/stl/test"
)

func TestTreeSet_Equal(t *testing.T) {
	hm1 := NewTreeSetWith[int](1)
	hm2 := NewTreeSetWith[int](1)
	hm3 := NewTreeSetWith[int](1, 2)
	stltest.AssertEq(t, hm1, hm2)
	stltest.AssertNotEq(t, hm1, hm3)
	stltest.AssertNotEq(t, hm2, hm3)
}
