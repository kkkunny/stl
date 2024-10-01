//go:build !goexperiment.rangefunc && !go1.23

package hashmap

type hashmapIter2[K, V any] interface{}
