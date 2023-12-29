package stlsync

import "sync"

// Locker 锁
type Locker interface {
	sync.Locker
	TryLock() bool
}
