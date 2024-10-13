package linkedhashmap

import (
	"encoding/json"
	"testing"

	"github.com/kkkunny/stl/container/tuple"
	stltest "github.com/kkkunny/stl/test"
	stlval "github.com/kkkunny/stl/value"
)

func TestAnyLinkedHashMap_Clone(t *testing.T) {
	hm1 := _NewAnyLinkedHashMapWith[int, int](1, 1, 2, 2)
	hm2 := stlval.Clone(hm1)
	stltest.AssertEq(t, hm1, hm2)
}

func TestAnyLinkedHashMap_Equal(t *testing.T) {
	hm1 := _NewAnyLinkedHashMapWith[int, int](1, 1, 2, 2)
	hm2 := _NewAnyLinkedHashMapWith[int, int](2, 2, 1, 1)
	hm3 := _NewAnyLinkedHashMapWith[int, int](1, 2, 2, 1)
	stltest.AssertNotEq(t, hm1, hm2)
	stltest.AssertNotEq(t, hm1, hm3)
	stltest.AssertNotEq(t, hm2, hm3)
}

func TestAnyLinkedHashMap_Json(t *testing.T) {
	hm := _NewAnyLinkedHashMapWith[string, int]("1", 1, "2", 2)
	data, err := json.Marshal(hm)
	if err != nil {
		panic(err)
	}
	stltest.AssertEq(t, string(data), "{\"1\":1,\"2\":2}")
}

func TestAnyLinkedHashMap_Length(t *testing.T) {
	hm := _NewAnyLinkedHashMapWith[int, int](1, 1, 2, 2)
	stltest.AssertEq(t, hm.Length(), 2)
}

func TestAnyLinkedHashMap_Set(t *testing.T) {
	hm := _NewAnyLinkedHashMap[int, int]()
	stltest.AssertEq(t, hm.Set(1, 1), 0)
	stltest.AssertEq(t, hm.Set(1, 2), 1)
	stltest.AssertEq(t, hm.Set(2, 2), 0)
	stltest.AssertEq(t, hm.Set(2, 3), 2)
}

func TestAnyLinkedHashMap_Get(t *testing.T) {
	hm := _NewAnyLinkedHashMapWith[int, int](1, 1, 2, 2)
	stltest.AssertEq(t, hm.Get(1, 2), 1)
	stltest.AssertEq(t, hm.Get(3), 0)
	stltest.AssertEq(t, hm.Get(3, 2), 2)
}

func TestAnyLinkedHashMap_ContainKey(t *testing.T) {
	hm := _NewAnyLinkedHashMapWith[int, int](1, 1, 2, 2)
	stltest.AssertEq(t, hm.Contain(1), true)
	stltest.AssertEq(t, hm.Contain(3), false)
}

func TestAnyLinkedHashMap_Remove(t *testing.T) {
	hm := _NewAnyLinkedHashMapWith[int, int](1, 1, 2, 2)
	stltest.AssertEq(t, hm.Remove(3), 0)
	stltest.AssertEq(t, hm.Remove(3, 3), 3)
	stltest.AssertEq(t, hm.Remove(1), 1)
}

func TestAnyLinkedHashMap_Clear(t *testing.T) {
	hm := _NewAnyLinkedHashMapWith[int, int](1, 1, 2, 2)
	stltest.AssertEq(t, hm.Length(), 2)
	hm.Clear()
	stltest.AssertEq(t, hm.Length(), 0)
}

func TestAnyLinkedHashMap_Empty(t *testing.T) {
	hm := _NewAnyLinkedHashMapWith[int, int](1, 1, 2, 2)
	stltest.AssertEq(t, hm.Empty(), false)
	hm.Clear()
	stltest.AssertEq(t, hm.Empty(), true)
}

func TestAnyLinkedHashMap_Keys(t *testing.T) {
	hm := _NewAnyLinkedHashMapWith[int, int](1, 1, 2, 2)
	stltest.AssertEq(t, hm.Keys(), []int{1, 2})
}

func TestAnyLinkedHashMap_Values(t *testing.T) {
	hm := _NewAnyLinkedHashMapWith[int, int](1, 1, 2, 2)
	stltest.AssertEq(t, hm.Values(), []int{1, 2})
}

func TestAnyLinkedHashMap_KeyValues(t *testing.T) {
	hm := _NewAnyLinkedHashMapWith[int, int](1, 1, 2, 2)
	stltest.AssertEq(t, hm.KeyValues(), []tuple.Tuple2[int, int]{tuple.Pack2(1, 1), tuple.Pack2(2, 2)})
}

func TestAnyLinkedHashMap_String(t *testing.T) {
	hm := _NewAnyLinkedHashMapWith[int, int](1, 1, 2, 2)
	stltest.AssertEq(t, hm.String(), "LinkedHashMap{1: 1, 2: 2}")
}
