package stltest

import (
	"testing"

	stlcmp "github.com/kkkunny/stl/cmp"
)

// AssertEq 必须相等
func AssertEq[T any](t *testing.T, lv, rv T) {
	if !stlcmp.Equal(lv, rv) {
		t.FailNow()
	}
}

// AssertNotEq 必须不相等
func AssertNotEq[T any](t *testing.T, lv, rv T) {
	if stlcmp.Equal(lv, rv) {
		t.FailNow()
	}
}
