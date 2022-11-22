package heap

import (
	"golang.org/x/exp/constraints"
)

// MaxHeap 大顶堆
type MaxHeap[T constraints.Ordered] struct {
	len      int
	data     []T
	hasValue []bool
}

// NewMaxHeap 新建大顶堆
func NewMaxHeap[T constraints.Ordered](e ...T) *MaxHeap[T] {
	h := new(MaxHeap[T])
	h.Push(e...)
	return h
}

// Length 获取长度 O(1)
func (self MaxHeap[T]) Length() int {
	return self.len
}

// Empty 是否为空 O(1)
func (self MaxHeap[T]) Empty() bool {
	return self.len == 0
}

// 压入调整
func (self *MaxHeap[T]) pushAdjust(i int) {
	for i > 0 {
		fi := (i - 1) / 2
		if self.data[i] > self.data[fi] {
			self.data[i], self.data[fi] = self.data[fi], self.data[i]
		}
		i = fi
	}
}

// Push 压入元素 O(logN)-O(NlogN)
func (self *MaxHeap[T]) Push(e ...T) {
loop:
	for _, elem := range e {
		self.len++
		for i := 0; i < len(self.hasValue); i++ {
			if !self.hasValue[i] {
				self.data[i] = elem
				self.hasValue[i] = true
				self.pushAdjust(i)
				continue loop
			}
		}
		self.data = append(self.data, elem)
		self.hasValue = append(self.hasValue, true)
		self.pushAdjust(self.len - 1)
	}
}

// 弹出调整
func (self *MaxHeap[T]) popAdjust(i int) {
	for {
		li, ri := 2*i+1, 2*i+2
		if li < len(self.data) && ri < len(self.data) && self.hasValue[li] && self.hasValue[ri] {
			self.hasValue[i] = true
			if self.data[li] >= self.data[ri] {
				self.data[i], self.data[li] = self.data[li], self.data[i]
				self.hasValue[li] = false
				i = li
			} else {
				self.data[i], self.data[ri] = self.data[ri], self.data[i]
				self.hasValue[ri] = false
				i = ri
			}
		} else if li < len(self.data) && self.hasValue[li] {
			self.data[i], self.data[li] = self.data[li], self.data[i]
			self.hasValue[i], self.hasValue[li] = true, false
			i = li
		} else if ri < len(self.data) && self.hasValue[ri] {
			self.data[i], self.data[ri] = self.data[ri], self.data[i]
			self.hasValue[i], self.hasValue[ri] = true, false
			i = ri
		} else {
			break
		}
	}
}

// Pop 弹出堆顶元素 O(NlogN)-O(N²logN)
func (self *MaxHeap[T]) Pop() T {
	top := self.data[0]
	self.hasValue[0] = false
	self.len--
	self.popAdjust(0)
	return top
}

// Peek 获取堆顶 O(1)
func (self MaxHeap[T]) Peek() T {
	return self.data[0]
}

// Clear 清空 O(1)
func (self *MaxHeap[T]) Clear() {
	if self.Empty() {
		return
	}
	self.data = nil
}

// Clone 克隆 O(N)
func (self MaxHeap[T]) Clone() *MaxHeap[T] {
	data := make([]T, len(self.data))
	copy(data, self.data)
	return &MaxHeap[T]{data: data}
}
