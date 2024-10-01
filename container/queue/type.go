package queue

// Queue 队列
type Queue[T any] []T

// NewQueue 新建队列
func NewQueue[T any]() Queue[T] {
	return make([]T, 0)
}

// NewQueueWith 新建指定元素的队列
func NewQueueWith[T any](vs ...T) Queue[T] {
	return vs
}
