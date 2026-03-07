package stlsync

import (
	"runtime"
	"sync/atomic"
)

const maxBackoff = 16

// 自旋
func spin(cond func() bool) {
	backoff := 1
	for cond() {
		for i := 0; i < backoff; i++ {
			runtime.Gosched()
		}
		if backoff < maxBackoff {
			backoff <<= 1
		}
	}
}

// SpinLock 自旋锁
// 避免线程上下文切换带来的开销，适用于锁竞争不激烈且锁定时间非常短的场景
type SpinLock struct {
	flag atomic.Bool
}

func NewSpinLock() Locker {
	return new(SpinLock)
}

func (lock *SpinLock) Lock() {
	spin(func() bool {
		return !lock.flag.CompareAndSwap(false, true)
	})
}

func (lock *SpinLock) TryLock() bool {
	return lock.flag.CompareAndSwap(false, true)
}

func (lock *SpinLock) Unlock() {
	lock.flag.CompareAndSwap(true, false)
}
