package queue

import "github.com/kkkunny/stl/container/dynarray"

// Queue 队列
type Queue[T any] dynarray.DynArray[T]

// NewQueue 新建队列
func NewQueue[T any]() Queue[T] {
	return Queue[T](dynarray.NewDynArray[T]())
}

// NewQueueWith 新建指定元素的队列
func NewQueueWith[T any](vs ...T) Queue[T] {
	return Queue[T](dynarray.NewDynArrayWith[T](vs...))
}
