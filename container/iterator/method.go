package iterator

func Collect[V any, Iter _Iter[V], Ctr _IterContainer[Ctr, V]](iter Iter) Ctr {
	iterator, ok := any(iter).(Iterator[V])
	if !ok{
		iterator = NewIterator(iter)
	}
	var ctr Ctr
	return ctr.NewWithIterator(iterator)
}

func Map[V1 any, Ctr1 _IterContainer[Ctr1, V1], V2 any, Ctr2 _IterContainer[Ctr2, V2]](from Ctr1, f func(V1) V2) Ctr2 {
	fromIter := from.Iterator()
	slice := make([]V2, fromIter.Length())
	var i int
	fromIter.Foreach(func(v V1) bool {
		slice[i] = f(v)
		i++
		return true
	})

	toIter := _NewSliceIterator(slice...)
	var ctr Ctr2
	return ctr.NewWithIterator(NewIterator(toIter))
}

func FlatMap[V1 any, Ctr1 _IterContainer[Ctr1, V1], V2 any, Ctr2 _IterContainer[Ctr2, V2]](from Ctr1, f func(V1) []V2) Ctr2 {
	fromIter := from.Iterator()
	slice := make([]V2, 0, fromIter.Length())
	fromIter.Foreach(func(v V1) bool {
		slice = append(slice, f(v)...)
		return true
	})

	toIter := _NewSliceIterator(slice...)
	var ctr Ctr2
	return ctr.NewWithIterator(NewIterator(toIter))
}

func All[V any, Ctr _IterContainer[Ctr, V]](ctr Ctr, f func(V) bool) bool {
	all := true
	ctr.Iterator().Foreach(func(v V) bool {
		all = all && f(v)
		return all
	})
	return all
}

func Any[V any, Ctr _IterContainer[Ctr, V]](ctr Ctr, f func(V) bool) bool {
	any := false
	ctr.Iterator().Foreach(func(v V) bool {
		any = any || f(v)
		return !any
	})
	return any
}

func Filter[V any, Ctr _IterContainer[Ctr, V]](ctr Ctr, f func(V) bool) Ctr {
	iter := ctr.Iterator()
	slice := make([]V, 0, iter.Length())
	iter.Foreach(func(v V) bool {
		if f(v) {
			slice = append(slice, v)
		}
		return true
	})
	return Collect[V, _Iter[V], Ctr](_NewSliceIterator[V](slice...))
}
