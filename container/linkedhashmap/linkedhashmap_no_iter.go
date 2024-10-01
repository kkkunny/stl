//go:build !goexperiment.rangefunc && !go1.23

package linkedhashmap

type linkedhashmapIter[K, V any] interface{}
