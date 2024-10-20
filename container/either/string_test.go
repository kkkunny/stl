package either

import (
	"testing"

	stltest "github.com/kkkunny/stl/test"
)

func TestEither_String(t *testing.T) {
	v := Left[int, uint](1)
	stltest.AssertEq(t, v.String(), "1")
}
