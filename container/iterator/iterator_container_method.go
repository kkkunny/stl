package iterator

// Foreach 遍历
func Foreach[T any](ctr IteratorContainer[T], f func(v T) bool) {
	IteratorForeach(ctr.Iterator(), f)
}

// Collect 收集
func Collect[V any, Ctr IteratorContainer[V]](iter Iterator[V]) Ctr {
	var ctr Ctr
	return ctr.NewWithIterator(iter).(Ctr)
}

// Map 映射
func Map[V1 any, V2 any, To IteratorContainer[V2]](from IteratorContainer[V1], f func(V1) V2) To {
	return Collect[V2, To](IteratorMap(from.Iterator(), f))
}

// FlatMap 扁平映射
func FlatMap[V1 any, V2 any, To IteratorContainer[V2]](from IteratorContainer[V1], f func(V1) []V2) To {
	return Collect[V2, To](IteratorFlatMap(from.Iterator(), f))
}

// All 所有元素都符合要求
func All[T any](ctr IteratorContainer[T], f func(T) bool) bool {
	return IteratorAll(ctr.Iterator(), f)
}

// Any 任意元素符合要求
func Any[T any](ctr IteratorContainer[T], f func(T) bool) bool {
	return IteratorAny(ctr.Iterator(), f)
}

// Filter 过滤
func Filter[T any, Ctr IteratorContainer[T]](ctr Ctr, f func(T) bool) Ctr {
	return Collect[T, Ctr](IteratorFilter(ctr.Iterator(), f))
}
