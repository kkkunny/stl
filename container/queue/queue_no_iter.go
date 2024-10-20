//go:build !goexperiment.rangefunc && !go1.23

package queue

type queueIter[T any] interface{}

func (self *_Queue[T]) NewWithIterator(iter stliter.Iterator[T]) any {
	data := make([]T, iter.Length())
	var i int
	for iter.Next() {
		data[i] = iter.Value()
		i++
	}
	return &_Queue[T]{data: data}
}

// Iterator 迭代器
func (self *_Queue[T]) Iterator() stliter.Iterator[T] {
	return stliter.NewSliceIterator(self.data...)
}
