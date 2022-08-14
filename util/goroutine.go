package util

import (
	"bytes"
	"runtime"
	"strconv"
)

// GetGoroutineID 获取协程id
func GetGoroutineID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}

// IsChanClosed 通道是否被关闭
func IsChanClosed[T any](ch chan T) bool {
	if ch == nil {
		return true
	}

	select {
	case _, ok := <-ch:
		return !ok
	default:
		return false
	}
}
