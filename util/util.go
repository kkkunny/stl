package util

// Ternary 三元运算
func Ternary[T any](c bool, t, f T) T {
	if c {
		return t
	} else {
		return f
	}
}
