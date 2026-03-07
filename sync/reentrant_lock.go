package stlsync

import (
	"sync"
	"sync/atomic"

	stlslices "github.com/kkkunny/stl/container/slices"
	stlruntime "github.com/kkkunny/stl/runtime"
)

// ReentrantLock 可重入锁
type ReentrantLock struct {
	locker    Locker
	owner     atomic.Int64
	recursion uint
}

func NewReentrantLock(locker ...Locker) *ReentrantLock {
	return &ReentrantLock{
		locker: stlslices.Last[Locker](locker, new(sync.Mutex)),
	}
}

func (l *ReentrantLock) Lock() {
	gid := int64(stlruntime.GetGoroutineID())
	if l.owner.Load() == gid {
		l.recursion++
		return
	}

	l.locker.Lock()
	l.owner.Store(gid)
	l.recursion = 1
}

func (l *ReentrantLock) TryLock() bool {
	gid := int64(stlruntime.GetGoroutineID())
	if l.owner.Load() == gid {
		l.recursion++
		return true
	}

	if !l.locker.TryLock() {
		return false
	}
	l.owner.Store(gid)
	l.recursion = 1
	return true
}

func (l *ReentrantLock) Unlock() {
	gid := int64(stlruntime.GetGoroutineID())
	if l.owner.Load() != gid {
		panic("unlock of wrong goroutine")
	}

	l.recursion--
	if l.recursion > 0 {
		return
	}

	l.owner.Store(-1)
	l.locker.Unlock()
}

// ReentrantRWLock 可重入读写锁
type ReentrantRWLock struct {
	locker        ReentrantLock
	rownersLocker RWLocker
	rowners       map[uint64]uint
}

func NewReentrantRWLock(locker ...RWLocker) *ReentrantRWLock {
	return &ReentrantRWLock{
		locker:        *NewReentrantLock(stlslices.Last[RWLocker](locker, new(sync.RWMutex))),
		rownersLocker: new(sync.RWMutex),
		rowners:       make(map[uint64]uint),
	}
}

func (l *ReentrantRWLock) Lock() {
	l.rownersLocker.Lock()
	l.rownersLocker.Unlock()
	l.locker.Lock()
}

func (l *ReentrantRWLock) TryLock() bool {
	l.rownersLocker.Lock()
	l.rownersLocker.Unlock()
	return l.locker.TryLock()
}

func (l *ReentrantRWLock) Unlock() {
	gid := stlruntime.GetGoroutineID()
	if l.locker.owner.Load() != int64(gid) {
		panic("unlock of wrong goroutine")
	}

	l.locker.recursion--
	if l.locker.recursion > 0 {
		return
	}

	l.rownersLocker.RLock()
	defer l.rownersLocker.RUnlock()
	rcount := l.rowners[gid]
	l.locker.owner.Store(-1)
	l.locker.locker.Unlock()

	// 锁降级 写锁 -> 读锁
	if rcount > 0 {
		l.locker.locker.(RWLocker).RLock()
	}
}

func (l *ReentrantRWLock) RLock() {
	gid := stlruntime.GetGoroutineID()
	l.rownersLocker.RLock()
	count := l.rowners[gid]
	l.rownersLocker.RUnlock()
	if count > 0 {
		l.rownersLocker.Lock()
		defer l.rownersLocker.Unlock()
		l.rowners[gid] = count + 1
		return
	}

	// 允许锁降级 写锁 -> 读锁
	if l.locker.owner.Load() == int64(gid) {
		l.rownersLocker.Lock()
		defer l.rownersLocker.Unlock()
		l.rowners[gid] = 1
		return
	}

	l.locker.locker.(RWLocker).RLock()
	l.rownersLocker.Lock()
	defer l.rownersLocker.Unlock()
	l.rowners[gid] = 1
}

func (l *ReentrantRWLock) TryRLock() bool {
	gid := stlruntime.GetGoroutineID()
	l.rownersLocker.RLock()
	count := l.rowners[gid]
	l.rownersLocker.RUnlock()
	if count > 0 {
		l.rownersLocker.Lock()
		defer l.rownersLocker.Unlock()
		l.rowners[gid] = count + 1
		return true
	}

	// 允许锁降级 写锁 -> 读锁
	if l.locker.owner.Load() == int64(gid) {
		l.rownersLocker.Lock()
		defer l.rownersLocker.Unlock()
		l.rowners[gid] = 1
		return true
	}

	if !l.locker.locker.(RWLocker).TryRLock() {
		return false
	}
	l.rownersLocker.Lock()
	defer l.rownersLocker.Unlock()
	l.rowners[gid] = 1
	return true
}

func (l *ReentrantRWLock) RUnlock() {
	gid := stlruntime.GetGoroutineID()
	l.rownersLocker.RLock()
	count := l.rowners[gid]
	l.rownersLocker.RUnlock()
	if count <= 0 {
		panic("sync: Unlock of wrong goroutine")
	}

	count--
	if count > 0 {
		l.rownersLocker.Lock()
		defer l.rownersLocker.Unlock()
		l.rowners[gid] = count
		return
	}

	l.rownersLocker.Lock()
	delete(l.rowners, gid)
	l.rownersLocker.Unlock()
	l.locker.locker.(RWLocker).RUnlock()
}
