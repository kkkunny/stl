package optional

import (
	"encoding/json"
	"fmt"
	"testing"
)

type TestMarshalJSONStruct struct {
	Field1 string
	Field2 Optional[string]
}

func TestOptional_MarshalJSON(t *testing.T) {
	data, err := json.Marshal(TestMarshalJSONStruct{
		Field1: "test",
		Field2: None[string](),
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}

func TestOptional_UnmarshalJSON(t *testing.T) {
	data := []byte(`{"Field1":"test","Field2":"test"}`)
	var v TestMarshalJSONStruct
	if err := json.Unmarshal(data, &v); err != nil {
		panic(err)
	}
	fmt.Println(v)
}
