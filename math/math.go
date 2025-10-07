package stlmath

import (
	"math"

	stlbasic "github.com/kkkunny/stl/basic"
)

// RoundTo 将m取整到n
func RoundTo[T stlbasic.Number](m, n T) T {
	return ((m + n - 1) / n) * n
}

// RoundToPowerOf 将m取整到n的幂
func RoundToPowerOf[T stlbasic.Number](m, n T) T {
	if n <= 1 {
		return 0
	}
	if m < 1 {
		return 1
	}
	log := math.Log(float64(m)) / math.Log(float64(n))
	nextPower := math.Ceil(log)
	return T(math.Pow(float64(n), nextPower))
}
