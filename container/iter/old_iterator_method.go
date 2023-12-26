package stliter

import stlbasic "github.com/kkkunny/stl/basic"

// IteratorForeach 遍历
func IteratorForeach[T any](iter Iterator[T], f func(v T) bool) {
	for iter.Next() {
		if !f(iter.Value()) {
			break
		}
	}
}

// IteratorMap 映射
func IteratorMap[V1, V2 any](from Iterator[V1], f func(V1) V2) Iterator[V2] {
	slice := make([]V2, from.Length())
	var i int
	IteratorForeach(from, func(v V1) bool {
		slice[i] = f(v)
		i++
		return true
	})
	return newSliceIterator(slice...)
}

// IteratorFlatMap 扁平映射
func IteratorFlatMap[V1, V2 any](from Iterator[V1], f func(V1) []V2) Iterator[V2] {
	slice := make([]V2, 0, from.Length())
	IteratorForeach(from, func(v V1) bool {
		slice = append(slice, f(v)...)
		return true
	})
	return newSliceIterator(slice...)
}

// IteratorAll 所有元素都符合要求
func IteratorAll[T any](iter Iterator[T], f func(T) bool) bool {
	all := true
	IteratorForeach(iter, func(v T) bool {
		all = all && f(v)
		return all
	})
	return all
}

// IteratorAny 任意元素符合要求
func IteratorAny[T any](iter Iterator[T], f func(T) bool) bool {
	any := false
	IteratorForeach(iter, func(v T) bool {
		any = any || f(v)
		return !any
	})
	return any
}

// IteratorFilter 过滤
func IteratorFilter[T any](iter Iterator[T], f func(T) bool) Iterator[T] {
	slice := make([]T, 0, iter.Length())
	IteratorForeach(iter, func(v T) bool {
		if f(v) {
			slice = append(slice, v)
		}
		return true
	})
	return newSliceIterator[T](slice...)
}

// IteratorContainAll 包含所有
func IteratorContainAll[T any](iter Iterator[T], v ...T) bool {
loop:
	for iter.Next() {
		for _, vv := range v {
			if stlbasic.Equal(iter.Value(), vv) {
				continue loop
			}
		}
		return false
	}
	return true
}

// IteratorContainAny 包含任意
func IteratorContainAny[T any](iter Iterator[T], v ...T) bool {
	for iter.Next() {
		for _, vv := range v {
			if stlbasic.Equal(iter.Value(), vv) {
				return true
			}
		}
	}
	return false
}

// IteratorContain 包含
func IteratorContain[T any](iter Iterator[T], v T) bool {
	return IteratorContainAny(iter, v)
}
