package hashset

import (
	"testing"

	stltest "github.com/kkkunny/stl/test"
)

func TestHashSet_Equal(t *testing.T) {
	hs1 := NewHashSetWith[int](1, 2, 2)
	hs2 := NewHashSetWith[int](1, 2, 1)
	hs3 := NewHashSetWith[int](1, 2, 3)
	stltest.AssertEq(t, hs1, hs2)
	stltest.AssertNotEq(t, hs1, hs3)
	stltest.AssertNotEq(t, hs2, hs3)
}
