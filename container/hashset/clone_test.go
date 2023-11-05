package hashset

import (
	"testing"

	stltest "github.com/kkkunny/stl/test"
)

func TestHashSet_Clone(t *testing.T) {
	hs1 := NewHashSetWith[int](1, 2, 2)
	hs2 := hs1.Clone()
	stltest.AssertEq(t, hs1, hs2)
}
