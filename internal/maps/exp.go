//go:build !go1.23

package maps

import "golang.org/x/exp/maps"

func Keys[M interface{ ~map[K]V }, K comparable, V any](m M) []K {
	return maps.Keys[M, K, V](m)
}
