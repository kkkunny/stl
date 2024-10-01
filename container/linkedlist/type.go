package linkedlist

type node[T any] struct {
	Prev, Next *node[T]
	Value      T
}

// LinkedList 链表
type LinkedList[T any] struct {
	length     uint
	root, tail *node[T]
}

func NewLinkedList[T any]() LinkedList[T] {
	return LinkedList[T]{}
}

func NewLinkedListWith[T any](vs ...T) LinkedList[T] {
	self := NewLinkedList[T]()
	for _, v := range vs {
		self.PushBack(v)
	}
	return self
}
