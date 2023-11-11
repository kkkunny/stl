package treeset

import (
	"testing"

	stltest "github.com/kkkunny/stl/test"
)

func TestTreeSet_Length(t *testing.T) {
	hm := NewTreeSetWith[int](1, 2)
	stltest.AssertEq(t, hm.Length(), 2)
}
