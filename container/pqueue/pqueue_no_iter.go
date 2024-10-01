//go:build !goexperiment.rangefunc && !go1.23

package pqueue

type pqueueIter[T any] interface{}
