package util

// 三元运算
func Ternary[T any](c bool, t, f func() T) T {
	if c {
		return t()
	} else {
		return f()
	}
}
