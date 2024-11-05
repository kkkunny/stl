//go:build !goexperiment.rangefunc && !go1.23

package queue

type queueIter[T any] interface{}
