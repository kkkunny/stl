package linkedlist

// Clone 克隆
func (self LinkedList[T]) Clone() LinkedList[T] {
	list := NewLinkedList[T]()
	for cursor := self.root; cursor != nil; cursor = cursor.Next {
		list.PushBack(cursor.Value)
	}
	return list
}
