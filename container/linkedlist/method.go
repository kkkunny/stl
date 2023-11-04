package linkedlist

func (self *LinkedList[T]) PushBack(v T) {
	node := &node[T]{Value: v}
	if self.root == nil {
		self.root = node
		self.tail = node
	} else {
		self.tail.Next = node
		node.Prev = self.tail
		self.tail = node
	}
	self.length++
}

func (self *LinkedList[T]) PushFront(v T) {
	node := &node[T]{Value: v}
	if self.root == nil {
		self.root = node
		self.tail = node
	} else {
		self.root.Prev = node
		node.Next = self.root
		self.root = node
	}
	self.length++
}

func (self *LinkedList[T]) PopBack() T {
	var v T
	if self.root == self.tail {
		v = self.root.Value
		self.root = nil
		self.tail = nil
	} else {
		v = self.tail.Value
		self.tail.Prev.Next = nil
		self.tail = self.tail.Prev
	}
	self.length--
	return v
}

func (self *LinkedList[T]) PopFront() T {
	var v T
	if self.root == self.tail {
		v = self.root.Value
		self.root = nil
		self.tail = nil
	} else {
		v = self.root.Value
		self.root.Next.Prev = nil
		self.root = self.root.Next
	}
	self.length--
	return v
}

func (self LinkedList[T]) Back() T {
	return self.tail.Value
}

func (self LinkedList[T]) Front() T {
	return self.root.Value
}

func (self *LinkedList[T]) Clear() {
	self.root = nil
	self.tail = nil
	self.length = 0
}

func (self LinkedList[T]) Empty() bool {
	return self.length == 0
}
