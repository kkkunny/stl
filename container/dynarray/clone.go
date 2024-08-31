package dynarray

import "golang.org/x/exp/slices"

// Clone 克隆
func (self DynArray[T]) Clone() DynArray[T] {
	self.init()
	newData := slices.Clone(*self.data)
	return DynArray[T]{data: &newData}
}
