//go:build !goexperiment.rangefunc && !go1.23

package stlmaps

type mapIter2[K, V any] interface {
}
