package hashmap

import (
	"encoding/json"
	"testing"

	stltest "github.com/kkkunny/stl/test"
)

func TestHashMap_Json(t *testing.T) {
	hm := NewHashMapWith[string, int]("1", 1, "2", 2)
	data, err := json.Marshal(hm)
	if err != nil {
		panic(err)
	}
	stltest.AssertEq(t, string(data), "{\"1\":1,\"2\":2}")
}
