package util

// Ternary 三元运算
func Ternary[T any](c bool, t, f T) T {
	if c {
		return t
	} else {
		return f
	}
}

// Must 必须没有异常
func Must(e error) {
	if e != nil {
		panic(e)
	}
}

// MustValue 必须没有异常并返回值
func MustValue[T any](v T, e error) T {
	if e != nil {
		panic(e)
	}
	return v
}
