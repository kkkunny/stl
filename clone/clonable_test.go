package clone

import (
	"testing"

	stltest "github.com/kkkunny/stl/test"
)

func TestClone_Slice(t *testing.T) {
	v := []int{1, 2, 3}
	c := Clone(v)
	stltest.AssertEq(t, c, v)
	stltest.AssertNotEq(t, &c[0], &v[0])
}

func TestClone_Array(t *testing.T) {
	v := [3]int{1, 2, 3}
	c := Clone(v)
	stltest.AssertEq(t, c, v)
}

func TestClone_Map(t *testing.T) {
	v := map[string]int{"1": 1, "2": 2}
	c := Clone(v)
	stltest.AssertEq(t, c, v)
}

func TestClone_Struct(t *testing.T) {
	v := struct{ A, B int }{1, 2}
	c := Clone(v)
	stltest.AssertEq(t, c, v)
}