package hashset

import (
	"testing"

	stltest "github.com/kkkunny/stl/test"
)

func TestHashSet_String(t *testing.T) {
	hm := NewHashSetWith[int](1, 2)
	stltest.AssertNotEq(t, hm.String(), "")
}
