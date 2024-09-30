package stlhash

import "github.com/kkkunny/maphash"

func GetMapHashFunc[T comparable]() func(v T) uint64 {
	hasher := maphash.NewHasher[T]()
	return func(vv T) uint64 {
		return hasher.Hash(vv)
	}
}

func MapHash[T comparable](v T) uint64 {
	return GetMapHashFunc[T]()(v)
}
