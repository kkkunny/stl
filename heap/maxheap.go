package heap

import (
	"stl/list"
	. "stl/types"
)

// 大顶堆
type MaxHeap[T Comparator[T]] struct {
	data *list.ArrayList[T]
}

// 新建大顶堆
func NewMaxHeap[T Comparator[T]](e ...T) *MaxHeap[T] {
	h := &MaxHeap[T]{
		data: list.NewArrayList[T](0, 0),
	}
	h.Push(e...)
	return h
}

// 转成字符串
func (self *MaxHeap[T]) String() string {
	return self.data.String()
}

// 获取长度
func (self *MaxHeap[T]) Length() Usize {
	return self.data.Length()
}

// 是否为空
func (self *MaxHeap[T]) Empty() bool {
	return self.data.Empty()
}

// 压入元素
func (self *MaxHeap[T]) Push(e ...T) {
	for _, v := range e {
		index := Isize(self.data.Length())
		self.data.Add(v)
		sIndex := index
		fIndex := (index - 1) / 2
		for sIndex != fIndex && fIndex >= 0 && self.data.Get(Usize(sIndex)).Compare(self.data.Get(Usize(fIndex))) > 0 {
			s, f := self.data.Get(Usize(sIndex)), self.data.Get(Usize(fIndex))
			self.data.Set(Usize(sIndex), f)
			self.data.Set(Usize(fIndex), s)
			sIndex = fIndex
			fIndex = (fIndex - 1) / 2
		}
	}
}

// 弹出堆顶元素
func (self *MaxHeap[T]) Pop() T {
	value := self.data.Get(0)
	lastIndex := self.data.Length() - 1
	last := self.data.Get(lastIndex)
	self.data.Set(0, last)
	self.data.Remove(lastIndex)
	var curIndex Usize
	length := lastIndex - 1
	for {
		leftIndex, rightIndex := 2*curIndex+1, 2*curIndex+2
		if leftIndex >= length { // 无左节点，无右节点
			break
		} else if rightIndex >= length { // 有左节点，无右节点
			if self.data.Get(curIndex).Compare(self.data.Get(leftIndex)) < 0 {
				f, l := self.data.Get(curIndex), self.data.Get(leftIndex)
				self.data.Set(curIndex, l)
				self.data.Set(leftIndex, f)
				curIndex = leftIndex
			} else {
				break
			}
		} else if self.data.Get(leftIndex).Compare(self.data.Get(rightIndex)) >= 0 { // 左节点比右节点小
			if self.data.Get(curIndex).Compare(self.data.Get(leftIndex)) < 0 {
				f, l := self.data.Get(curIndex), self.data.Get(leftIndex)
				self.data.Set(curIndex, l)
				self.data.Set(leftIndex, f)
				curIndex = leftIndex
			} else {
				break
			}
		} else { // 左节点比右节点大
			if self.data.Get(curIndex).Compare(self.data.Get(rightIndex)) < 0 {
				f, r := self.data.Get(curIndex), self.data.Get(rightIndex)
				self.data.Set(curIndex, r)
				self.data.Set(rightIndex, f)
				curIndex = rightIndex
			} else {
				break
			}
		}
	}
	return value
}

// 提前获取堆顶
func (self *MaxHeap[T]) Peek() T {
	return self.data.Get(0)
}

// 清空
func (self *MaxHeap[T]) Clear() {
	self.data.Clear()
}

// 克隆
func (self *MaxHeap[T]) Clone() *MaxHeap[T] {
	return &MaxHeap[T]{
		data: self.data.Clone(),
	}
}

// 获取迭代器
func (self *MaxHeap[T]) Iterator() *list.ArrayListIterator[T] {
	return self.data.Begin()
}
