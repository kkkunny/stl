package stlbasic

// Ternary 三目运算
func Ternary[T any](cond bool, t, f T) T {
	if cond {
		return t
	}
	return f
}

// TernaryAction 三目运算行为
func TernaryAction[T any](cond bool, t, f func() T) T {
	if cond {
		return t()
	}
	return f()
}
