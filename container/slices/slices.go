package stlslices

import (
	"math/rand"
	"slices"
	"time"

	stlbasic "github.com/kkkunny/stl/basic"
)

func Map[T, F any](slice []T, f func(i int, e T) F) []F {
	res := make([]F, len(slice))
	for i, e := range slice {
		res[i] = f(i, e)
	}
	return res
}

func MapError[T, F any](slice []T, f func(i int, e T) (F, error)) (res []F, err error) {
	res = make([]F, len(slice))
	for i, e := range slice {
		res[i], err = f(i, e)
		if err != nil {
			return nil, err
		}
	}
	return res, nil
}

func FlatMap[T, F any](slice []T, f func(i int, e T) []F) []F {
	res := make([]F, 0, len(slice))
	for i, e := range slice {
		res = append(res, f(i, e)...)
	}
	return res
}

func FlatMapError[T, F any](slice []T, f func(i int, e T) ([]F, error)) (res []F, err error) {
	res = make([]F, 0, len(slice))
	var es []F
	for i, e := range slice {
		es, err = f(i, e)
		if err != nil {
			return nil, err
		}
		res = append(res, es...)
	}
	return res, nil
}

func All[T any, TS ~[]T](slice TS, f func(i int, e T) bool) bool {
	for i, e := range slice {
		if !f(i, e) {
			return false
		}
	}
	return true
}

func Any[T any, TS ~[]T](slice TS, f func(i int, e T) bool) bool {
	for i, e := range slice {
		if f(i, e) {
			return true
		}
	}
	return false
}

func Filter[T any, TS ~[]T](slice TS, f func(i int, e T) bool) TS {
	res := make(TS, 0, len(slice))
	for i, e := range slice {
		if f(i, e) {
			res = append(res, e)
		}
	}
	return res
}

func ContainAll[T any, TS ~[]T](slice TS, v ...T) bool {
loop:
	for _, e := range slice {
		for _, vv := range v {
			if stlbasic.Equal(e, vv) {
				continue loop
			}
		}
		return false
	}
	return true
}

func ContainAny[T any, TS ~[]T](slice TS, v ...T) bool {
	for _, e := range slice {
		for _, vv := range v {
			if stlbasic.Equal(e, vv) {
				return true
			}
		}
	}
	return false
}

func Contain[T any, TS ~[]T](slice TS, v T) bool {
	return ContainAny(slice, v)
}

func Sort[T any, TS ~[]T](slice TS, reverse ...bool) TS {
	slice = slices.Clone(slice)
	slices.SortFunc(slice, func(l, r T) int {
		if len(reverse) > 0 && reverse[0] {
			return stlbasic.Order(r, l)
		} else {
			return stlbasic.Order(l, r)
		}
	})
	return slice
}

func As[T any, TS ~[]T, F any, FS ~[]F](slice TS) FS {
	return Map(slice, func(_ int, e T) F {
		return any(e).(F)
	})
}

// DiffTo 返回l中r没有的值
func DiffTo[T any, TS ~[]T](l, r TS) (res TS) {
	res = make([]T, 0, len(l))
loop:
	for _, le := range l {
		for _, re := range r {
			if stlbasic.Equal(le, re) {
				continue loop
			}
		}
		res = append(res, le)
	}
	return res
}

// Diff 返回l和r中各在对方没有的值
func Diff[T any, TS ~[]T](l, r TS) TS {
	return Union(DiffTo(l, r), DiffTo(r, l))
}

// Union 联合
func Union[T any, TS ~[]T](l, r TS) TS {
	return append(l, r...)
}

// Intersect 返回l和r中各在对方有的值
func Intersect[T any, TS ~[]T](l, r TS) (res TS) {
	return DiffTo(l, DiffTo(l, r))
}

func RemoveRepeat[T any, TS ~[]T](slice TS) (res TS) {
	if len(slice) <= 1 {
		return slice
	}

	conflictMap := make([]bool, len(slice))
	res = make(TS, 0, len(slice))
	for i, ie := range slice {
		if conflictMap[i] {
			continue
		}
		for j, je := range slice[i+1:] {
			if !stlbasic.Equal(ie, je) {
				continue
			}
			conflictMap[i+j+1] = true
			if !conflictMap[i] {
				conflictMap[i] = true
				res = append(res, ie)
			}
		}
		if !conflictMap[i] {
			res = append(res, ie)
		}
	}
	return res
}

func First[T any, TS ~[]T](slice TS, defaultValue ...T) (v T) {
	if Empty(slice) {
		return Last(defaultValue, v)
	}
	return slice[0]
}

func Last[T any, TS ~[]T](slice TS, defaultValue ...T) (v T) {
	if Empty(slice) {
		return Last(defaultValue, v)
	}
	return slice[len(slice)-1]
}

func Empty[T any, TS ~[]T](slice TS) bool {
	return len(slice) == 0
}

// And 同Diff
func And[T any, TS ~[]T](l, r TS) TS {
	return Diff(l, r)
}

// Or 同Union+RemoveRepeat
func Or[T any, TS ~[]T](l, r TS) TS {
	return RemoveRepeat(Union(l, r))
}

func Shuffle[T any, TS ~[]T](slice TS) TS {
	res := make(TS, len(slice))
	copy(res, slice)
	rand.New(rand.NewSource(time.Now().UnixNano())).Shuffle(len(slice), func(i, j int) {
		res[i], res[j] = res[j], res[i]
	})
	return res
}

func ToMap[T any, TS ~[]T, K comparable, V any, KV ~map[K]V](slice TS, mapFn func(T) (K, V)) KV {
	res := make(KV, len(slice))
	for _, e := range slice {
		k, v := mapFn(e)
		res[k] = v
	}
	return res
}
