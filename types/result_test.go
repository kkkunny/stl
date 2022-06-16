package types

import (
	"errors"
	"github.com/kkkunny/stl/util"
	"testing"
)

func TestResult(t *testing.T) {
	a := Ok(1)
	util.AssertEq(t, a.IsOk(), true)
	util.AssertEq(t, a.IsErr(), false)
	util.AssertEq(t, a.Ok().Unwrap(), 1)
	util.AssertEq(t, a.Err().IsNone(), true)
	util.AssertEq(t, a.Unwrap(), 1)
	a = Err[int](errors.New(""))
	util.AssertEq(t, a.UnwrapOr(1), 1)
	util.AssertEq(t, a.UnwrapOrDefault(), 0)
}
