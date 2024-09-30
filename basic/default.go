package stlbasic

// Defaultable 有默认值的
type Defaultable[Self any] interface {
	Default() Self
}

// Default 获取默认值
func Default[T any]() T {
	var self T
	if tmp, ok := any(self).(Defaultable[T]); ok {
		return tmp.Default()
	} else {
		return self
	}
}
