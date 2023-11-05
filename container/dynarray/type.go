package dynarray

const initialCapacity uint = 5 // 初始容量

// DynArray 动态数组
type DynArray[T any] struct {
	data *[]T
}

// NewDynArray 新建动态数组
func NewDynArray[T any]() DynArray[T] {
	return NewDynArrayWithCapacity[T](initialCapacity)
}

// NewDynArrayWithCapacity 新建指定容量的动态数组
func NewDynArrayWithCapacity[T any](cap uint) DynArray[T] {
	data := make([]T, 0, cap)
	return DynArray[T]{data: &data}
}

// NewDynArrayWithLength 新建指定长度的动态数组
func NewDynArrayWithLength[T any](l uint) DynArray[T] {
	data := make([]T, l)
	return DynArray[T]{data: &data}
}

// NewDynArrayWith 新建指定元素的动态数组
func NewDynArrayWith[T any](vs ...T) DynArray[T] {
	return DynArray[T]{data: &vs}
}
