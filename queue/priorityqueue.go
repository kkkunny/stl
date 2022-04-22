package queue

import (
	"fmt"
	"github.com/kkkunny/stl/types"
	"golang.org/x/exp/constraints"
)

// 优先级队列
type PriorityQueue[P constraints.Ordered, V any] struct {
	data []types.Pair[P, V]
}

// 新建优先级队列
func NewPriorityQueue[P constraints.Ordered, V any]() *PriorityQueue[P, V] {
	return new(PriorityQueue[P, V])
}

// 转成字符串 O(N)
func (self *PriorityQueue[P, V]) String() string {
	return fmt.Sprintf("%v", self.data)
}

// 获取长度 O(1)
func (self *PriorityQueue[P, V]) Length() int {
	return len(self.data)
}

// 是否为空 O(1)
func (self *PriorityQueue[P, V]) Empty() bool {
	return len(self.data) == 0
}

// 压入队列 O(1)
func (self *PriorityQueue[P, V]) Push(p P, v V) {
	self.data = append(self.data, types.NewPair(p, v))
	sIndex, fIndex := len(self.data), (len(self.data)-1)/2
	for sIndex != fIndex && fIndex >= 0 && self.data[sIndex].First > self.data[fIndex].First {
		s, f := self.data[sIndex], self.data[fIndex]
		self.data[sIndex], self.data[fIndex] = f, s
		sIndex, fIndex = fIndex, (fIndex-1)/2
	}
}

// 弹出队列 O(1)
func (self *PriorityQueue[P, V]) Pop() (P, V) {
	value := self.data[0]
	if len(self.data) == 1 {
		self.data = self.data[1:]
		return value.First, value.Second
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
			if self.data[curIndex].First < self.data[leftIndex].First {
				f, l := self.data[curIndex], self.data[leftIndex]
				self.data[curIndex] = l
				self.data[leftIndex] = f
				curIndex = leftIndex
			} else {
				break
			}
		} else if self.data[leftIndex].First >= self.data[rightIndex].First { // 左节点比右节点小
			if self.data[curIndex].First < self.data[leftIndex].First {
				f, l := self.data[curIndex], self.data[leftIndex]
				self.data[curIndex] = l
				self.data[leftIndex] = f
				curIndex = leftIndex
			} else {
				break
			}
		} else { // 左节点比右节点大
			if self.data[curIndex].First < self.data[rightIndex].First {
				f, r := self.data[curIndex], self.data[rightIndex]
				self.data[curIndex] = r
				self.data[rightIndex] = f
				curIndex = rightIndex
			} else {
				break
			}
		}
	}
	return value.First, value.Second
}

// 获取队首 O(1)
func (self *PriorityQueue[P, V]) Peek() (P, V) {
	return self.data[0].First, self.data[0].Second
}

// 清空 O(1)
func (self *PriorityQueue[P, V]) Clear() {
	if self.Empty() {
		return
	}
	self.data = nil
}

// 克隆 O(N)
func (self *PriorityQueue[P, V]) Clone() *PriorityQueue[P, V] {
	data := make([]types.Pair[P, V], len(self.data))
	copy(data, self.data)
	return &PriorityQueue[P, V]{data: data}
}
