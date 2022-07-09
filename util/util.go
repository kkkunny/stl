package util

import (
	"bytes"
	"encoding/json"
	"runtime"
	"strconv"
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
func Must(e error) {
	if e != nil {
		panic(e)
	}
}

// MustValue 必须没有异常并返回值
func MustValue[T any](v T, e error) T {
	if e != nil {
		panic(e)
	}
	return v
}

// GetGoroutineID 获取协程id
func GetGoroutineID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}
