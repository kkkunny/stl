package clone

import (
	"github.com/cespare/xxhash/v2"
	"github.com/elastic/go-freelru"
)

type retType struct {
	f  func(v any) any
	ok bool
}

var cloneFuncCache = func() *freelru.ShardedLRU[string, retType] {
	cache, err := freelru.NewSharded[string, retType](20, func(s string) uint32 { return uint32(xxhash.Sum64String(s)) })
	if err != nil {
		panic(err)
	}
	return cache
}()
