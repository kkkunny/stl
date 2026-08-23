package stlruntime

import _ "unsafe"

// runtime 内部的 CPU 级自旋（执行 PAUSE 指令 N 次）
//
//go:linkname procyield runtime.procyield
func procyield(cycles uint32)

// ProcYield CPU级自旋，不让出 CPU
func ProcYield(cycles uint32) {
	procyield(cycles)
	// // 无法直接发 PAUSE 指令，用多次原子读模拟 CPU 空转
	// // 编译器不会优化掉 atomic 操作
	// var sink int32
	// for i := 0; i < 30; i++ {
	// 	atomic.LoadInt32(&sink)
	// }
}

// 让出 goroutine 调度，但当前 goroutine 留在本地 P 的 runq 上（比 Gosched 轻量）
//
//go:linkname goyield runtime.goyield
func goyield()

// GoYield 让出当前 goroutine 的时间片，goroutine 留在本地 P 的 runq 中，很快会被重新调度；
// 比 runtime.Gosched 轻量（不进全局 runq）
func GoYield() {
	goyield()
}
