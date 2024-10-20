//go:build goexperiment.rangefunc || go1.23

package stliter

import (
	"iter"
	"math/rand"
	"slices"
	"time"

	stlcmp "github.com/kkkunny/stl/cmp"
	iter2 "github.com/kkkunny/stl/internal/iter"
	"github.com/kkkunny/stl/internal/reflect"
	stlval "github.com/kkkunny/stl/value"
)

func FromSlice[T any](slice []T) iter.Seq[T] {
	return func(yield func(T) bool) {
		for _, e := range slice {
			if !yield(e) {
				return
			}
		}
	}
}

func Collect[Elem, To any](seq iter.Seq[Elem]) To {
	switch any(stlval.Default[To]()).(type) {
	case []Elem:
		return any(slices.Collect(seq)).(To)
	default:
		to := reflect.TypeFor[To]()
		fnObj, ok := iter2.ContainerFunctions[to.String()]
		if ok {
			fn := fnObj.(func(seq iter.Seq[Elem]) To)
			return fn(seq)
		}
		panic("unreachable")
	}
}

func Map[T, F any](seq iter.Seq[T], f func(i int, e T) F) iter.Seq[F] {
	return func(yield func(F) bool) {
		var i int
		for e := range seq {
			if !yield(f(i, e)) {
				return
			}
			i++
		}
	}
}

func FlatMap[T, F any](seq iter.Seq[T], f func(i int, e T) []F) iter.Seq[F] {
	return func(yield func(F) bool) {
		var i int
		for e := range seq {
			for _, f := range f(i, e) {
				if !yield(f) {
					return
				}
				i++
			}
		}
	}
}

func All[T any](seq iter.Seq[T], f func(i int, e T) bool) bool {
	var i int
	for e := range seq {
		if !f(i, e) {
			return false
		}
		i++
	}
	return true
}

func Any[T any](seq iter.Seq[T], f func(i int, e T) bool) bool {
	var i int
	for e := range seq {
		if f(i, e) {
			return true
		}
		i++
	}
	return false
}

func Filter[T any](seq iter.Seq[T], f func(i int, e T) bool) iter.Seq[T] {
	return func(yield func(T) bool) {
		var i int
		for e := range seq {
			if !f(i, e) {
				continue
			}
			if !yield(e) {
				return
			}
			i++
		}
	}
}

func ContainAll[T any](seq iter.Seq[T], v ...T) bool {
	equalFn := stlcmp.GetEqualFunc[T]()
loop:
	for _, vv := range v {
		for e := range seq {
			if !equalFn(vv, e) {
				continue loop
			}
			return false
		}
	}
	return true
}

func ContainAny[T any](seq iter.Seq[T], v ...T) bool {
	equalFn := stlcmp.GetEqualFunc[T]()
	for _, vv := range v {
		for e := range seq {
			if equalFn(vv, e) {
				return true
			}
		}
	}
	return false
}

func Contain[T any](seq iter.Seq[T], v T) bool {
	return ContainAny(seq, v)
}

func Sort[T any](seq iter.Seq[T], reverse ...bool) iter.Seq[T] {
	var needReverse bool
	if len(reverse) > 0 {
		needReverse = reverse[len(reverse)-1]
	}
	cmpFn := stlcmp.GetCompareFunc[T]()
	res := slices.SortedStableFunc(seq, func(l, r T) int {
		if needReverse {
			return -cmpFn(l, r)
		} else {
			return cmpFn(l, r)
		}
	})
	return FromSlice(res)
}

func Shuffle[T any](seq iter.Seq[T]) iter.Seq[T] {
	slice := Collect[T, []T](seq)
	rand.New(rand.NewSource(time.Now().UnixNano())).Shuffle(len(slice), func(i, j int) {
		slice[i], slice[j] = slice[j], slice[i]
	})
	return FromSlice(slice)
}

func ToSeq2[T, K, V any](seq iter.Seq[T], mapFn func(i int, e T) (K, V)) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		var i int
		for e := range seq {
			if !yield(mapFn(i, e)) {
				return
			}
			i++
		}
	}
}

func Random[T any](seq iter.Seq[T]) T {
	slice := Collect[T, []T](seq)
	index := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(len(slice))
	return slice[index]
}
