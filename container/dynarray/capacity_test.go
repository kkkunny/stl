package dynarray

import (
	"testing"

	stltest "github.com/kkkunny/stl/test"
)

func TestDynArray_Capacity(t *testing.T) {
	var v DynArray[int]
	stltest.AssertEq(t, v.Capacity(), 0)
	v = NewDynArrayWithCapacity[int](10)
	stltest.AssertEq(t, v.Capacity(), 10)
}
