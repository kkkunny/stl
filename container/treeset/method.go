package treeset

import (
	"github.com/kkkunny/stl/container/treemap"
)

// Add 插入值
func (self *TreeSet[T]) Add(v T) bool {
	exist := (*treemap.TreeMap[T, struct{}])(self).ContainKey(v)
	if exist {
		return false
	}
	(*treemap.TreeMap[T, struct{}])(self).Set(v, struct{}{})
	return true
}

// Contain 是否包含值
func (self TreeSet[T]) Contain(v T) bool {
	return treemap.TreeMap[T, struct{}](self).ContainKey(v)
}

// Remove 移除值
func (self *TreeSet[T]) Remove(v T) bool {
	exist := (*treemap.TreeMap[T, struct{}])(self).ContainKey(v)
	if !exist {
		return false
	}
	(*treemap.TreeMap[T, struct{}])(self).Remove(v)
	return true
}

// Clear 清空
func (self *TreeSet[T]) Clear() {
	(*treemap.TreeMap[T, struct{}])(self).Clear()
}

// Empty 是否为空
func (self TreeSet[T]) Empty() bool {
	return treemap.TreeMap[T, struct{}](self).Empty()
}

// ToSlice 转成切片
func (self TreeSet[T]) ToSlice() []T {
	return treemap.TreeMap[T, struct{}](self).Keys()
}
