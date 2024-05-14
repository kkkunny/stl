package linkedhashmap

import (
	"encoding/json"
	"testing"

	stltest "github.com/kkkunny/stl/test"
)

func TestLinkedHashMap_Json(t *testing.T) {
	hm := NewLinkedHashMapWith[string, int]("1", 1, "2", 2)
	data, err := json.Marshal(hm)
	if err != nil {
		panic(err)
	}
	stltest.AssertEq(t, string(data), "{\"1\":1,\"2\":2}")
}
