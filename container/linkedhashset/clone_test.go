package linkedhashset

import (
	"testing"

	stltest "github.com/kkkunny/stl/test"
)

func TestLinkedHashSet_Clone(t *testing.T) {
	hm1 := NewLinkedHashSetWith[int](1, 2)
	hm2 := hm1.Clone()
	stltest.AssertEq(t, hm1, hm2)
}
