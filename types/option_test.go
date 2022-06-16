package types

import (
	"github.com/kkkunny/stl/util"
	"testing"
)

func TestOption(t *testing.T) {
	a := Some(1)
	util.AssertEq(t, a.IsSome(), true)
	util.AssertEq(t, a.IsNone(), false)
	util.AssertEq(t, a.Unwrap(), 1)
	a = None[int]()
	util.AssertEq(t, a.UnwrapOr(1), 1)
	util.AssertEq(t, a.UnwrapOrDefault(), 0)
}
