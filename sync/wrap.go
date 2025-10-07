package stlsync

type Wrap[T any] struct {
	Locker
	elem T
}

func NewWrap[T any](locker Locker, v T) *Wrap[T] {
	return &Wrap[T]{
		Locker: locker,
		elem:   v,
	}
}

func (wrap *Wrap[T]) Set(v T) {
	wrap.Lock()
	defer wrap.Unlock()
	wrap.elem = v
}

func (wrap *Wrap[T]) Get() T {
	rwLocker, ok := wrap.Locker.(RWLocker)
	if !ok {
		wrap.Lock()
		defer wrap.Unlock()
	} else {
		rwLocker.RLock()
		defer rwLocker.RUnlock()
	}
	return wrap.elem
}

func (wrap *Wrap[T]) Swap(v T) T {
	wrap.Lock()
	defer wrap.Unlock()
	vv := wrap.elem
	wrap.elem = v
	return vv
}

func (wrap *Wrap[T]) WrapFn(fn func(v T) T) T {
	wrap.Lock()
	defer wrap.Unlock()
	wrap.elem = fn(wrap.elem)
	return wrap.elem
}
