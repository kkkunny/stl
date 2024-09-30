package stlslices

import (
	"math/rand"
	"time"

	"golang.org/x/exp/slices"

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

func All[T any](slice []T, f func(i int, e T) bool) bool {
	for i, e := range slice {
		if !f(i, e) {
			return false
		}
	}
	return true
}

func Any[T any](slice []T, f func(i int, e T) bool) bool {
	for i, e := range slice {
		if f(i, e) {
			return true
		}
	}
	return false
}

func Filter[T any](slice []T, filter func(i int, e T) bool) []T {
	res := make([]T, 0, len(slice))
	for i, e := range slice {
		if filter(i, e) {
			res = append(res, e)
		}
	}
	return res
}

func ContainAll[T any](slice []T, v ...T) bool {
loop:
	for _, vv := range v {
		for _, e := range slice {
			if stlbasic.Equal(e, vv) {
				continue loop
			}
		}
		return false
	}
	return true
}

func ContainAny[T any](slice []T, v ...T) bool {
	for _, vv := range v {
		for _, e := range slice {
			if stlbasic.Equal(e, vv) {
				return true
			}
		}
	}
	return false
}

func Contain[T any](slice []T, v T) bool {
	return ContainAny(slice, v)
}

func Sort[T any](slice []T, reverse ...bool) []T {
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

func As[T any, F any](slice []T) []F {
	return Map(slice, func(_ int, e T) F {
		return any(e).(F)
	})
}

// DiffTo 返回l中r没有的值
func DiffTo[T any](l, r []T) (res []T) {
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
func Diff[T any](l, r []T) []T {
	return Union(DiffTo(l, r), DiffTo(r, l))
}

// Union 联合
func Union[T any](l, r []T) []T {
	return append(l, r...)
}

// Intersect 返回l和r中各在对方有的值
func Intersect[T any](l, r []T) (res []T) {
	return DiffTo(l, DiffTo(l, r))
}

func RemoveRepeat[T any](slice []T) (res []T) {
	if len(slice) <= 1 {
		return slice
	}

	conflictMap := make([]bool, len(slice))
	res = make([]T, 0, len(slice))
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

func First[T any](slice []T, defaultValue ...T) (v T) {
	if Empty(slice) {
		return Last(defaultValue, v)
	}
	return slice[0]
}

func Last[T any](slice []T, defaultValue ...T) (v T) {
	if Empty(slice) {
		return Last(defaultValue, v)
	}
	return slice[len(slice)-1]
}

func Empty[T any](slice []T) bool {
	return len(slice) == 0
}

// And 同Diff
func And[T any](l, r []T) []T {
	return Diff(l, r)
}

// Or 同Union+RemoveRepeat
func Or[T any](l, r []T) []T {
	return RemoveRepeat(Union(l, r))
}

func Shuffle[T any](slice []T) []T {
	res := make([]T, len(slice))
	copy(res, slice)
	rand.New(rand.NewSource(time.Now().UnixNano())).Shuffle(len(slice), func(i, j int) {
		res[i], res[j] = res[j], res[i]
	})
	return res
}

func ToMap[T any, K comparable, V any](slice []T, mapFn func(T) (K, V)) map[K]V {
	res := make(map[K]V, len(slice))
	for _, e := range slice {
		k, v := mapFn(e)
		res[k] = v
	}
	return res
}

func Random[T any](slice []T) T {
	index := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(len(slice))
	return slice[index]
}

func FindFirst[T any](slice []T, filter func(i int, e T) bool, defaultValue ...T) (T, bool) {
	for i, e := range slice {
		if filter(i, e) {
			return e, true
		}
	}
	return First(defaultValue), false
}

func FindLast[T any](slice []T, filter func(i int, e T) bool, defaultValue ...T) (T, bool) {
	if !Empty(slice) {
		for i := len(slice) - 1; i >= 0; i-- {
			if filter(i, slice[i]) {
				return slice[i], true
			}
		}
	}
	return Last(defaultValue), false
}

func Clone[T any](slice []T) []T {
	newSlice := make([]T, len(slice), cap(slice))
	copy(newSlice, slice)
	return newSlice
}

func Repeat[T any](v T, n int) []T {
	slice := make([]T, n)
	for i := 0; i < n; i++ {
		slice[i] = v
	}
	return slice
}

// Equal 比较两个切片里的元素是否相等
func Equal[T any](l, r []T) bool {
	if len(l) != len(r) {
		return false
	}
	for i, lv := range l {
		rv := r[i]
		if !stlbasic.Equal(lv, rv) {
			return false
		}
	}
	return true
}
