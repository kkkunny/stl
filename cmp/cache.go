package stlcmp

import (
	"github.com/cespare/xxhash/v2"
	"github.com/elastic/go-freelru"
)

type cmpRetType struct {
	f  func(l, r any) int
	ok bool
}

var cmpFuncCache = func() *freelru.ShardedLRU[string, cmpRetType] {
	cache, err := freelru.NewSharded[string, cmpRetType](20, func(s string) uint32 { return uint32(xxhash.Sum64String(s)) })
	if err != nil {
		panic(err)
	}
	return cache
}()

type eqRetType struct {
	f  func(l, r any) bool
	ok bool
}

var eqFuncCache = func() *freelru.ShardedLRU[string, eqRetType] {
	cache, err := freelru.NewSharded[string, eqRetType](20, func(s string) uint32 { return uint32(xxhash.Sum64String(s)) })
	if err != nil {
		panic(err)
	}
	return cache
}()
