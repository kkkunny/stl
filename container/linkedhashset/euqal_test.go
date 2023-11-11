package linkedhashset

import (
	"testing"

	stltest "github.com/kkkunny/stl/test"
)

func TestLinkedHashSet_Equal(t *testing.T) {
	hm1 := NewLinkedHashSetWith[int](1)
	hm2 := NewLinkedHashSetWith[int](1)
	hm3 := NewLinkedHashSetWith[int](1, 2)
	stltest.AssertEq(t, hm1, hm2)
	stltest.AssertNotEq(t, hm1, hm3)
	stltest.AssertNotEq(t, hm2, hm3)
}
