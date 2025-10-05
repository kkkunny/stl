package stlhash

import (
	"reflect"

	"github.com/kkkunny/maphash"
)

// Hashable 可哈希的
type Hashable interface {
	Hash() uint64
}

// GetHashFunc 获取哈希函数，若没有会panic
func GetHashFunc[T any]() func(v T) uint64 {
	t := reflect.TypeFor[T]()
	switch {
	case t.Implements(reflect.TypeFor[Hashable]()):
		return func(vv T) uint64 {
			return any(vv).(Hashable).Hash()
		}
	default:
		hasher := maphash.NewHasher2[T]()
		return func(vv T) uint64 {
			return hasher.Hash(vv)
		}
	}
}

// Hash 获取哈希，若没有会panic
func Hash[T any](v T) uint64 {
	return GetHashFunc[T]()(v)
}
