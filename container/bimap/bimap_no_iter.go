//go:build !goexperiment.rangefunc && !go1.23

package bimap

type bimapIter[T, E any] interface{}
