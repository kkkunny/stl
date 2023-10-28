package stlbasic

// Ternary 三目运算
func Ternary[T any](cond bool, t, f T) T {
	if cond {
		return t
	}
	return f
}
