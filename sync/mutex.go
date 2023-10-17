package stlsync

import "sync"

type Mutex[T any] struct {
	sync.Mutex
	value T
}

func NewMutex[T any](v T) Mutex[T] {
	return Mutex[T]{value: v}
}

func (self Mutex[T]) GetValue() T {
	self.Lock()
	defer self.Unlock()
	return self.UnsafeGetValue()
}

func (self Mutex[T]) SetValue(v T) T {
	self.Lock()
	defer self.Unlock()
	return self.SetValue(v)
}

func (self Mutex[T]) UnsafeGetValue() T {
	return self.value
}

func (self Mutex[T]) UnsafeSetValue(v T) T {
	src := self.value
	self.value = v
	return src
}
