package stlsync

// RWLocker 读写锁
type RWLocker interface {
	Locker
	RLock()
	RUnlock()
	TryRLock() bool
}
