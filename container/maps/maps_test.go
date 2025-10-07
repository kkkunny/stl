package stlmaps

import (
	"testing"

	stltest "github.com/kkkunny/stl/test"
)

func TestContainKey(t *testing.T) {
	hmap := map[int]int{
		1: 1,
		2: 2,
	}
	stltest.AssertEq(t, ContainKey(hmap, 1), true)
	stltest.AssertEq(t, ContainKey(hmap, 2), true)
	stltest.AssertEq(t, ContainKey(hmap, 3), false)
}

func TestContainAnyKeys(t *testing.T) {
	hmap := map[int]int{
		1: 1,
		2: 2,
	}
	stltest.AssertEq(t, ContainAnyKeys(hmap, 1, 2), true)
	stltest.AssertEq(t, ContainAnyKeys(hmap, 3), false)
}

func TestContainAllKeys(t *testing.T) {
	hmap := map[int]int{
		1: 1,
		2: 2,
	}
	stltest.AssertEq(t, ContainAllKeys(hmap, 1, 2), true)
	stltest.AssertEq(t, ContainAllKeys(hmap, 1, 3), false)
}
