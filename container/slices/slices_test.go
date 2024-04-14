package stlslices

import (
	"testing"

	stltest "github.com/kkkunny/stl/test"
)

func TestDiffTo(t *testing.T) {
	stltest.AssertEq(t, DiffTo([]int{1, 2}, []int{2, 3}), []int{1})
}

func TestDiff(t *testing.T) {
	stltest.AssertEq(t, Diff([]int{1, 2}, []int{2, 3}), []int{1, 3})
}

func TestIntersect(t *testing.T) {
	stltest.AssertEq(t, Intersect([]int{1, 2}, []int{2, 3}), []int{2})
}

func TestRemoveRepeat(t *testing.T) {
	stltest.AssertEq(t, RemoveRepeat([]int{1, 2, 3, 1, 2, 3, 1, 2}), []int{1, 2, 3})
}
