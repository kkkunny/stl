package stltest

import (
	"testing"

	stlbasic "github.com/kkkunny/stl/basic"
)

// AssertEq 必须相等
func AssertEq[T any](t *testing.T, lv, rv T) {
	if !stlbasic.Equal(lv, rv) {
		t.FailNow()
	}
}

// AssertNotEq 必须不相等
func AssertNotEq[T any](t *testing.T, lv, rv T) {
	if stlbasic.Equal(lv, rv) {
		t.FailNow()
	}
}
