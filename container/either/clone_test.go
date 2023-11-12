package either

import (
	"testing"

	stltest "github.com/kkkunny/stl/test"
)

func TestEither_Clone(t *testing.T) {
	v := Left[int, uint](1)
	stltest.AssertEq(t, v, v.Clone())
}
