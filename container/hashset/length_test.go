package hashset

import (
	"testing"

	stltest "github.com/kkkunny/stl/test"
)

func TestHashSet_Length(t *testing.T) {
	hm := NewHashSetWith[int](1, 2)
	stltest.AssertEq(t, hm.Length(), 2)
}
