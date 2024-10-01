//go:build go1.23

package maps

import (
	"maps"
)

func Keys[Map interface{ ~map[K]V }, K comparable, V any](m Map) []K {
	keys := make([]K, len(m))
	var i int
	for k := range maps.Keys[Map, K, V](m) {
		keys[i] = k
		i++
	}
	return keys
}
