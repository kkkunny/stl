//go:build !goexperiment.rangefunc && !go1.23

package treemap

type treemapIter[K, V any] interface{}
