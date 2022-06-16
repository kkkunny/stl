package util

import "testing"

// AssertEq 断言相等
func AssertEq[T comparable](t *testing.T, v1, v2 T) bool {
	if v1 != v2 {
		t.FailNow()
		return false
	}
	return true
}

// AssertNeq 断言不相等
func AssertNeq[T comparable](t *testing.T, v1, v2 T) bool {
	if v1 == v2 {
		t.FailNow()
		return false
	}
	return true
}
