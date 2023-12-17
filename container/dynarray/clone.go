package dynarray

import "slices"

// Clone 克隆
func (self DynArray[T]) Clone() DynArray[T] {
	self.init()
	newData := slices.Clone(*self.data)
	return DynArray[T]{data: &newData}
}
