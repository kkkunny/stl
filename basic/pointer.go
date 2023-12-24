package stlbasic

// Ptr 获取值指针
func Ptr[T any](v T) *T {
	return &v
}
