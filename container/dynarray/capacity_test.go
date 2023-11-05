package dynarray

import (
	"testing"

	stltest "github.com/kkkunny/stl/test"
)

func TestDynArray_Capacity(t *testing.T) {
	var v DynArray[int]
	stltest.AssertEq(t, v.Capacity(), initialCapacity)
	v = NewDynArrayWithCapacity[int](20)
	stltest.AssertEq(t, v.Capacity(), 20)
}
