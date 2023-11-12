package either

import (
	"testing"

	stltest "github.com/kkkunny/stl/test"
)

func TestEither_Equal(t *testing.T) {
	v1 := Left[int, uint](1)
	v2 := Left[int, uint](1)
	v3 := Left[int, uint](2)
	stltest.AssertEq(t, v1, v2)
	stltest.AssertNotEq(t, v1, v3)
	stltest.AssertNotEq(t, v2, v3)
}
