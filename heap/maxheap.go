package heap

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// 大顶堆
type MaxHeap[T constraints.Ordered] struct {
	data []T
}

// 新建大顶堆
func NewMaxHeap[T constraints.Ordered](e ...T) *MaxHeap[T] {
	h := new(MaxHeap[T])
	h.Push(e...)
	return h
}

// 转成字符串 O(N)
func (self *MaxHeap[T]) String() string {
	return fmt.Sprintf("%v", self.data)
}

// 获取长度 O(1)
func (self *MaxHeap[T]) Length() int {
	return len(self.data)
}

// 是否为空 O(1)
func (self *MaxHeap[T]) Empty() bool {
	return len(self.data) == 0
}

// 压入元素 O(logN)-O(NlogN)
func (self *MaxHeap[T]) Push(e ...T) {
	for _, v := range e {
		index := len(self.data)
		self.data = append(self.data, v)
		sIndex := index
		fIndex := (index - 1) / 2
		for sIndex != fIndex && fIndex >= 0 && self.data[sIndex] > self.data[fIndex] {
			s, f := self.data[sIndex], self.data[fIndex]
			self.data[sIndex] = f
			self.data[fIndex] = s
			sIndex = fIndex
			fIndex = (fIndex - 1) / 2
		}
	}
}

// 弹出堆顶元素 O(NlogN)-O(N²logN)
func (self *MaxHeap[T]) Pop() T {
	value := self.data[0]
	if len(self.data) == 1 {
		self.data = self.data[1:]
		return value
	}
	lastIndex := len(self.data) - 1
	last := self.data[lastIndex]
	self.data[0] = last
	self.data = self.data[:len(self.data)-1]
	var curIndex int
	length := lastIndex - 1
	for {
		leftIndex, rightIndex := 2*curIndex+1, 2*curIndex+2
		if leftIndex >= length { // 无左节点，无右节点
			break
		} else if rightIndex >= length { // 有左节点，无右节点
			if self.data[curIndex] < self.data[leftIndex] {
				f, l := self.data[curIndex], self.data[leftIndex]
				self.data[curIndex] = l
				self.data[leftIndex] = f
				curIndex = leftIndex
			} else {
				break
			}
		} else if self.data[leftIndex] >= self.data[rightIndex] { // 左节点比右节点小
			if self.data[curIndex] < self.data[leftIndex] {
				f, l := self.data[curIndex], self.data[leftIndex]
				self.data[curIndex] = l
				self.data[leftIndex] = f
				curIndex = leftIndex
			} else {
				break
			}
		} else { // 左节点比右节点大
			if self.data[curIndex] < self.data[rightIndex] {
				f, r := self.data[curIndex], self.data[rightIndex]
				self.data[curIndex] = r
				self.data[rightIndex] = f
				curIndex = rightIndex
			} else {
				break
			}
		}
	}
	return value
}

// 提前获取堆顶 O(1)
func (self *MaxHeap[T]) Peek() T {
	return self.data[0]
}

// 清空 O(1)
func (self *MaxHeap[T]) Clear() {
	if self.Empty() {
		return
	}
	self.data = nil
}

// 克隆 O(N)
func (self *MaxHeap[T]) Clone() *MaxHeap[T] {
	data := make([]T, len(self.data))
	copy(data, self.data)
	return &MaxHeap[T]{data: data}
}
