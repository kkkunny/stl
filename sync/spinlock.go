package stlsync

import (
	"runtime"
	"sync/atomic"

	stlruntime "github.com/kkkunny/stl/runtime"
)

const (
	activeSpinCount  = 4  // 纯自旋轮次
	activeSpinSize   = 30 // 每轮执行 PAUSE 指令次数
	passiveSpinCount = 2  // osyield 轮次
)

// 自旋
func spin(cond func() bool) {
	for iter := 0; cond(); iter++ {
		if iter < activeSpinCount {
			// ====== 第1级：纯自旋 ======
			// 执行 CPU PAUSE 指令，真正的忙等
			// 不让出任何东西，延迟最低
			stlruntime.ProcYield(activeSpinSize)
		} else if iter < activeSpinCount+passiveSpinCount {
			// ====== 第2级：OS 线程级让出 ======
			// 线程让出时间片但仍在 run queue，很快会被重新调度
			// 比 Gosched 轻量：不涉及 goroutine 调度
			stlruntime.OsYield()
		} else {
			// ====== 第3级：goroutine 级让出 ======
			// 最后的兜底，防止长时间拿不到锁时饿死其他 goroutine
			runtime.Gosched()
		}
	}
}

// SpinLock 自旋锁
// 避免线程上下文切换带来的开销，适用于锁竞争不激烈且锁定时间非常短的场景
type SpinLock struct {
	flag atomic.Int32
}

func NewSpinLock() *SpinLock {
	return &SpinLock{}
}

func (l *SpinLock) Lock() {
	spin(func() bool {
		return !l.flag.CompareAndSwap(0, -1)
	})
}

func (l *SpinLock) TryLock() bool {
	return l.flag.CompareAndSwap(0, -1)
}

func (l *SpinLock) Unlock() {
	l.flag.CompareAndSwap(-1, 0)
}

// SpinRWLock 自旋读写锁
// 避免线程上下文切换带来的开销，适用于锁竞争不激烈且锁定时间非常短的场景
type SpinRWLock struct {
	locker SpinLock
}

func NewSpinRWLock() *SpinRWLock {
	return &SpinRWLock{locker: *NewSpinLock()}
}

func (l *SpinRWLock) Lock() {
	l.locker.Lock()
}

func (l *SpinRWLock) TryLock() bool {
	return l.locker.TryLock()
}

func (l *SpinRWLock) Unlock() {
	l.locker.Unlock()
}

func (l *SpinRWLock) RLock() {
	spin(func() bool {
		v := l.locker.flag.Load()
		if v < 0 {
			return true
		}
		return !l.locker.flag.CompareAndSwap(v, v+1)
	})
}

func (l *SpinRWLock) TryRLock() bool {
	for {
		v := l.locker.flag.Load()
		if v < 0 {
			return false
		}
		if l.locker.flag.CompareAndSwap(v, v+1) {
			return true
		}
	}
}

func (l *SpinRWLock) RUnlock() {
	v := l.locker.flag.Load()
	if v <= 0 {
		panic("sync: RUnlock of unlocked RWLock")
	}
	l.locker.flag.Add(-1)
}
