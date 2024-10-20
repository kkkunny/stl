package stlos

import (
	"testing"

	stltest "github.com/kkkunny/stl/test"
)

func TestExist(t *testing.T) {
	exist, err := Exist("111")
	if err != nil {
		panic(err)
	}
	stltest.AssertEq(t, exist, false)
	exist, err = Exist("file_test.go")
	if err != nil {
		panic(err)
	}
	stltest.AssertEq(t, exist, true)
}
