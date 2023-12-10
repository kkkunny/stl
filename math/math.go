package stlmath

import (
	stlbasic "github.com/kkkunny/stl/basic"
)

// RoundTo 将m取整到n
func RoundTo[T stlbasic.Number](m, n T) T {
	return ((m + n - 1) / n) * n
}
