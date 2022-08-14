package json

import (
	"bytes"
	"encoding/json"
	"strings"
)

// Json json序列化
func Json(v any) ([]byte, error) {
	var res bytes.Buffer
	encoder := json.NewEncoder(&res)
	err := encoder.Encode(v)
	return res.Bytes(), err
}

// JsonToString json序列化
func JsonToString(v any) (string, error) {
	var res bytes.Buffer
	encoder := json.NewEncoder(&res)
	err := encoder.Encode(v)
	return res.String(), err
}

// UnJson json反序列化
func UnJson(v any, data []byte) error {
	decoder := json.NewDecoder(bytes.NewReader(data))
	return decoder.Decode(v)
}

// UnJsonFromString json反序列化
func UnJsonFromString(v any, data string) error {
	decoder := json.NewDecoder(strings.NewReader(data))
	return decoder.Decode(v)
}
