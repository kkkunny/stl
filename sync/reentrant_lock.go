package sync

import (
	"github.com/kkkunny/stl/util"
	"sync"
)

// ReentrantLock 可重入锁
type ReentrantLock struct {
	mux sync.Mutex
	c   chan struct{}
	n   int
	gid uint64
}

// Lock 加锁
func (self *ReentrantLock) Lock() {
	gid := util.GetGoroutineID()
	for {
		self.mux.Lock()
		if self.c == nil {
			self.c = make(chan struct{}, 1)
		}
		if self.n == 0 || self.gid == gid {
			self.n++
			self.gid = gid
			self.mux.Unlock()
			break
		}
		self.mux.Unlock()
		<-self.c
	}
}

// TryLock 尝试加锁
func (self *ReentrantLock) TryLock() bool {
	gid := util.GetGoroutineID()
	self.mux.Lock()
	defer self.mux.Unlock()
	if self.c == nil {
		self.c = make(chan struct{}, 1)
	}
	if self.n == 0 || self.gid == gid {
		self.n++
		self.gid = gid
		return true
	}
	return false
}

// Unlock 解锁
func (self *ReentrantLock) Unlock() {
	gid := util.GetGoroutineID()
	self.mux.Lock()
	if self.n <= 0 || gid != self.gid {
		self.mux.Unlock()
		panic("sync: unlock of unlocked mutex")
	}
	self.n--
	self.mux.Unlock()
	select {
	case self.c <- struct{}{}:
	default:
	}
}
