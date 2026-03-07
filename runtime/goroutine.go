package stlruntime

import (
	"bytes"
	"runtime"
	"strconv"
)

// GetGoroutineID 获取当前goroutine的id
func GetGoroutineID() uint64 {
	var buf [64]byte
	var s = buf[:runtime.Stack(buf[:], false)]
	s = s[len("goroutine "):]
	s = s[:bytes.IndexByte(s, ' ')]
	gid, _ := strconv.ParseUint(string(s), 10, 64)
	return gid
}
