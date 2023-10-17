package stlsync

import "sync"

type RWMutex[T any] struct {
	sync.RWMutex
	value T
}

func NewRWMutex[T any](v T) Mutex[T] {
	return Mutex[T]{value: v}
}

func (self RWMutex[T]) GetValue() T {
	self.RLock()
	defer self.RUnlock()
	return self.UnsafeGetValue()
}

func (self RWMutex[T]) SetValue(v T) T {
	self.Lock()
	defer self.Unlock()
	return self.SetValue(v)
}

func (self RWMutex[T]) UnsafeGetValue() T {
	return self.value
}

func (self RWMutex[T]) UnsafeSetValue(v T) T {
	src := self.value
	self.value = v
	return src
}
