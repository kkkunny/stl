package btree

import stlbasic "github.com/kkkunny/stl/basic"

// 二叉树节点
type BTreeNode[T any] struct {
	Father      *BTreeNode[T]
	Value       T
	Left, Right *BTreeNode[T]
}

// 最小的大数
func (self *BTreeNode[T]) GetSmallestLargeNode()*BTreeNode[T]{
	if self.Right == nil{
		return self
	}
	return self.Right.GetLeftestNode()
}

// 最大的小数
func (self *BTreeNode[T]) GetLargestSmallNode()*BTreeNode[T]{
	if self.Left == nil{
		return self
	}
	return self.Left.GetRightestNode()
}

func (self *BTreeNode[T]) GetLeftestNode()*BTreeNode[T]{
	if self.Left == nil{
		return self
	}
	return self.Left.GetLeftestNode()
}

func (self *BTreeNode[T]) GetRightestNode()*BTreeNode[T]{
	if self.Right == nil{
		return self
	}
	return self.Right.GetRightestNode()
}

func (self *BTreeNode[T]) IsLeft()bool{
	return self.Father != nil && self.Father.Left == self
}

func (self *BTreeNode[T]) IsRight()bool{
	return self.Father != nil && self.Father.Right == self
}

func (self *BTreeNode[T]) SetLeft(n *BTreeNode[T]){
	if n == nil{
		return
	}
	self.Left = n
	n.Father = self
}

func (self *BTreeNode[T]) SetRight(n *BTreeNode[T]){
	if n == nil{
		return
	}
	self.Right = n
	n.Father = self
}

// 二叉树
type BTree[T any] struct {
	length uint
	root *BTreeNode[T]
}

func NewBTree[T any]()BTree[T]{
	return BTree[T]{}
}

func (self BTree[T]) Length()uint{
	return self.length
}

func (self *BTree[T]) PushNode(node *BTreeNode[T])*BTreeNode[T]{
	if self.root == nil{
		self.root = node
		node.Father, node.Left, node.Right = nil, nil, nil
		self.length++
		return node
	}else{
		cursor := self.root
		for {
			order := stlbasic.Order[T](node.Value, cursor.Value)
			switch{
			case order < 0:
				if cursor.Left == nil{
					node.Left, node.Right = nil, nil
					cursor.SetLeft(node)
					self.length++
					return node
				}else{
					cursor = cursor.Left
				}
			case order == 0:
				return cursor
			case order > 0:
				if cursor.Right == nil{
					node.Left, node.Right = nil, nil
					cursor.SetRight(node)
					self.length++
					return node
				}else{
					cursor = cursor.Right
				}
			}
		}
	}
}

func (self *BTree[T]) Push(v T) *BTreeNode[T] {
	return self.PushNode(&BTreeNode[T]{Value: v})
}

func (self BTree[T]) Top()*BTreeNode[T]{
	return self.root
}

func (self *BTree[T]) RemoveNode(node *BTreeNode[T])*BTreeNode[T]{
	if self.length == 1{
		self.root = nil
		self.length--
		node.Left, node.Right = nil, nil
		return node
	}else{
		newNode := node.GetSmallestLargeNode()
		if newNode == node{
			newNode = node.GetLargestSmallNode()
		}
		if newNode == node{
			self.length--
			if node.IsLeft(){
				node.Father.Left = nil
			}else if node.IsRight(){
				node.Father.Right = nil
			}
			node.Father = nil
			return node
		}else{
			newNode = self.RemoveNode(newNode)
			newNode.SetLeft(node.Left)
			newNode.SetRight(node.Right)
			newNode.Father = node.Father
			if self.root == node{
				self.root = newNode
			}
			node.Father, node.Left, node.Right = nil, nil, nil
			return node
		}
	}
}

func (self BTree[T]) Find(v T)*BTreeNode[T]{
	cursor := self.root
loop:
	for cursor!=nil{
		order := stlbasic.Order(v, cursor.Value)
		switch {
		case order < 0:
			cursor = cursor.Left
		case order == 0:
			break loop
		case order > 0:
			cursor = cursor.Right
		}
	}
	return cursor
}

func (self *BTree[T]) Remove(v T)*BTreeNode[T]{
	node := self.Find(v)
	if node == nil{
		return nil
	}
	return self.RemoveNode(node)
}
