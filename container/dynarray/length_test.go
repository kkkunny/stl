package dynarray

import (
	"testing"

	stltest "github.com/kkkunny/stl/test"
)

func TestDynArray_Length(t *testing.T) {
	v := NewDynArrayWithLength[int](10)
	stltest.AssertEq(t, v.Capacity(), 10)
	stltest.AssertEq(t, v.Length(), 10)
}
