package lazy

import (
	"sync"

	stlsync "github.com/kkkunny/stl/sync"
	stlval "github.com/kkkunny/stl/value"
)

// Getter 懒加载
func Getter[T comparable](fn func() (T, error)) func() (T, error) {
	getter := GetterWith[struct{}, T](func(_ struct{}) (T, error) {
		return fn()
	})
	return func() (T, error) {
		return getter(struct{}{})
	}
}

// GetterWith 带入参的懒加载，不会根据入参不同做不同缓存，所以入参最好选择无状态的
func GetterWith[I, O comparable](fn func(in I) (O, error)) func(in I) (O, error) {
	var lock stlsync.RWLocker = new(sync.RWMutex)
	var cache O
	zeroV := stlval.Default[O]()
	return func(in I) (O, error) {
		v := func() O {
			lock.RLock()
			defer lock.RUnlock()
			return cache
		}()
		if v != zeroV {
			return v, nil
		}

		lock.Lock()
		defer lock.Unlock()
		if cache != zeroV {
			return cache, nil
		}

		v, err := fn(in)
		if err != nil {
			return zeroV, err
		}
		cache = v
		return v, nil
	}
}
