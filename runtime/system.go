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

// 系统调用让出 CPU 时间片（sched_yield / SwitchToThread）
//
//go:linkname osyield runtime.osyield
func osyield()

// OsYield 让出时间片，但线程仍可运行
func OsYield() {
	osyield()
}
