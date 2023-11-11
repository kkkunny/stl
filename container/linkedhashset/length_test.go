package linkedhashset

import (
	"testing"

	stltest "github.com/kkkunny/stl/test"
)

func TestLinkedHashSet_Length(t *testing.T) {
	hm := NewLinkedHashSetWith[int](1, 2)
	stltest.AssertEq(t, hm.Length(), 2)
}
