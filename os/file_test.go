package stlos

import (
	"testing"

	stlerror "github.com/kkkunny/stl/error"
	stltest "github.com/kkkunny/stl/test"
)

func TestExist(t *testing.T) {
	stltest.AssertEq(t, stlerror.MustWith(Exist("111")), false)
	stltest.AssertEq(t, stlerror.MustWith(Exist("file_test.go")), true)
}
