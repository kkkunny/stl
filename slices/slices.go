package stlslices

import (
	"slices"

	stlbasic "github.com/kkkunny/stl/basic"
)

func Map[T, F any](slice []T, f func(i int, e T) F) []F {
	res := make([]F, len(slice))
	for i, e := range slice {
		res[i] = f(i, e)
	}
	return res
}

func FlatMap[T, F any](slice []T, f func(i int, e T) []F) []F {
	res := make([]F, 0, len(slice))
	for i, e := range slice {
		res = append(res, f(i, e)...)
	}
	return res
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
