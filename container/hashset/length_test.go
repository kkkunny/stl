package hashset

import (
	"testing"

	stltest "github.com/kkkunny/stl/test"
)

func TestHashSet_Length(t *testing.T) {
	hs := NewHashSetWith[int](1, 2, 2)
	stltest.AssertEq(t, hs.Length(), 2)
}
