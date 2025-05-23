package linkedhashmap

import (
	"encoding/json"
	"testing"

	"github.com/kkkunny/stl/clone"
	"github.com/kkkunny/stl/container/tuple"
	stltest "github.com/kkkunny/stl/test"
)

func TestStdLinkedHashMap_Clone(t *testing.T) {
	hm1 := _NewStdLinkedHashMapWith[int, int](1, 1, 2, 2)
	hm2 := clone.Clone(hm1)
	stltest.AssertEq(t, hm1, hm2)
}

func TestStdLinkedHashMap_Equal(t *testing.T) {
	hm1 := _NewStdLinkedHashMapWith[int, int](1, 1, 2, 2)
	hm2 := _NewStdLinkedHashMapWith[int, int](2, 2, 1, 1)
	hm3 := _NewStdLinkedHashMapWith[int, int](1, 2, 2, 1)
	stltest.AssertNotEq(t, hm1, hm2)
	stltest.AssertNotEq(t, hm1, hm3)
	stltest.AssertNotEq(t, hm2, hm3)
}

func TestStdLinkedHashMap_Json(t *testing.T) {
	hm := _NewStdLinkedHashMapWith[string, int]("1", 1, "2", 2)
	data, err := json.Marshal(hm)
	if err != nil {
		panic(err)
	}
	stltest.AssertEq(t, string(data), "{\"1\":1,\"2\":2}")
}

func TestStdLinkedHashMap_Length(t *testing.T) {
	hm := _NewStdLinkedHashMapWith[int, int](1, 1, 2, 2)
	stltest.AssertEq(t, hm.Length(), 2)
}

func TestStdLinkedHashMap_Set(t *testing.T) {
	hm := _NewStdLinkedHashMap[int, int]()
	stltest.AssertEq(t, hm.Set(1, 1), 0)
	stltest.AssertEq(t, hm.Set(1, 2), 1)
	stltest.AssertEq(t, hm.Set(2, 2), 0)
	stltest.AssertEq(t, hm.Set(2, 3), 2)
}

func TestStdLinkedHashMap_Get(t *testing.T) {
	hm := _NewStdLinkedHashMapWith[int, int](1, 1, 2, 2)
	stltest.AssertEq(t, hm.Get(1, 2), 1)
	stltest.AssertEq(t, hm.Get(3), 0)
	stltest.AssertEq(t, hm.Get(3, 2), 2)
}

func TestStdLinkedHashMap_ContainKey(t *testing.T) {
	hm := _NewStdLinkedHashMapWith[int, int](1, 1, 2, 2)
	stltest.AssertEq(t, hm.Contain(1), true)
	stltest.AssertEq(t, hm.Contain(3), false)
}

func TestStdLinkedHashMap_Remove(t *testing.T) {
	hm := _NewStdLinkedHashMapWith[int, int](1, 1, 2, 2)
	stltest.AssertEq(t, hm.Remove(3), 0)
	stltest.AssertEq(t, hm.Remove(3, 3), 3)
	stltest.AssertEq(t, hm.Remove(1), 1)
}

func TestStdLinkedHashMap_Clear(t *testing.T) {
	hm := _NewStdLinkedHashMapWith[int, int](1, 1, 2, 2)
	stltest.AssertEq(t, hm.Length(), 2)
	hm.Clear()
	stltest.AssertEq(t, hm.Length(), 0)
}

func TestStdLinkedHashMap_Empty(t *testing.T) {
	hm := _NewStdLinkedHashMapWith[int, int](1, 1, 2, 2)
	stltest.AssertEq(t, hm.Empty(), false)
	hm.Clear()
	stltest.AssertEq(t, hm.Empty(), true)
}

func TestStdLinkedHashMap_Keys(t *testing.T) {
	hm := _NewStdLinkedHashMapWith[int, int](1, 1, 2, 2)
	stltest.AssertEq(t, hm.Keys(), []int{1, 2})
}

func TestStdLinkedHashMap_Values(t *testing.T) {
	hm := _NewStdLinkedHashMapWith[int, int](1, 1, 2, 2)
	stltest.AssertEq(t, hm.Values(), []int{1, 2})
}

func TestStdLinkedHashMap_KeyValues(t *testing.T) {
	hm := _NewStdLinkedHashMapWith[int, int](1, 1, 2, 2)
	stltest.AssertEq(t, hm.KeyValues(), []tuple.Tuple2[int, int]{tuple.Pack2(1, 1), tuple.Pack2(2, 2)})
}

func TestStdLinkedHashMap_String(t *testing.T) {
	hm := _NewStdLinkedHashMapWith[int, int](1, 1, 2, 2)
	stltest.AssertEq(t, hm.String(), "LinkedHashMap{1: 1, 2: 2}")
}

func TestStdLinkedHashMap_MarshalJSON(t *testing.T) {
	hm := _NewStdLinkedHashMapWith[string, int]("1", 1)
	data, err := json.Marshal(hm)
	stltest.AssertEq(t, err, nil)
	stltest.AssertNotEq(t, string(data), "{\"1\": 1}")
}

func TestStdLinkedHashMap_UnmarshalJSON(t *testing.T) {
	hm1 := _NewStdLinkedHashMapWith[string, int]("1", 1)
	data, err := json.Marshal(hm1)
	stltest.AssertEq(t, err, nil)
	hm2 := _NewStdLinkedHashMapWith[string, int]()
	err = json.Unmarshal(data, hm2)
	stltest.AssertEq(t, hm1, hm2)
}
