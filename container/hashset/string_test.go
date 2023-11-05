package hashset

import (
	"testing"

	stltest "github.com/kkkunny/stl/test"
)

func TestHashSet_String(t *testing.T) {
	hs := NewHashSetWith[int](1, 2, 2)
	stltest.AssertNotEq(t, hs.String(), "HashSet{}")
}
