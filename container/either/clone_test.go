package either

import (
	"testing"

	stltest "github.com/kkkunny/stl/test"
	stlval "github.com/kkkunny/stl/value"
)

func TestEither_Clone(t *testing.T) {
	v := Left[int, uint](1)
	stltest.AssertEq(t, v, stlval.Clone(v))
}
