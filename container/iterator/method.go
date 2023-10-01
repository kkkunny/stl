package iterator

func Map[C1, V1 any, Ctr1 _IterContainer[C1, V1], C2, V2 any, Ctr2 _IterContainer[C2, V2]](from Ctr1, f func(V1)V2) C2 {
	fromIter := from.Iterator()
	slice := make([]V2, fromIter.Length())
	var i int
	fromIter.Foreach(func(v V1) {
		slice[i] = f(v)
		i++
	})

	toIter := _NewSliceIterator(slice...)
	var Ctr Ctr2
	return Ctr.NewWithIterator(NewIterator(toIter))
}
