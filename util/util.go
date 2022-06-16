package util

import (
	"bytes"
	"encoding/json"
)

// Ternary 三元运算
func Ternary[T any](c bool, t, f T) T {
	if c {
		return t
	} else {
		return f
	}
}

// Json json序列化
func Json(v any) ([]byte, error) {
	var res bytes.Buffer
	encoder := json.NewEncoder(&res)
	err := encoder.Encode(v)
	return res.Bytes(), err
}

// UnJson json反序列化
func UnJson(v any, data []byte) error {
	decoder := json.NewDecoder(bytes.NewReader(data))
	return decoder.Decode(v)
}

// Must 必须没有异常
func Must[T any](v T, e error) T {
	if e != nil {
		panic(e)
	}
	return v
}
