package stlbasic

// Is 类型是否是
func Is[T any](v any) bool {
	_, ok := v.(T)
	return ok
}
