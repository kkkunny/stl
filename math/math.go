package math

import (
	"github.com/kkkunny/stl/types"
	"golang.org/x/exp/constraints"
	"math"
)

// Abs 获得绝对值
func Abs[T types.Number](v T) T {
	if v < 0 {
		return -v
	} else {
		return v
	}
}

// Max 获取最大值
func Max[T constraints.Ordered](v ...T) (max T) {
	for i, vv := range v {
		if i == 0 || vv > max {
			max = vv
		}
	}
	return max
}

// Min 获取最小值
func Min[T constraints.Ordered](v ...T) (min T) {
	for i, vv := range v {
		if i == 0 || vv < min {
			min = vv
		}
	}
	return min
}

// FloatMod 浮点数取余/模
func FloatMod[T constraints.Float](l, r T) T {
	if r == 0 || math.IsInf(float64(l), 0) || math.IsNaN(float64(l)) || math.IsNaN(float64(r)) {
		return T(math.NaN())
	}

	ll, rr := Abs(l), Abs(r)
	if ll < rr {
		return l
	} else if ll == rr {
		return 0
	}

	for ll > rr {
		ll = ll - rr
	}
	if l < 0 {
		return -ll
	} else {
		return ll
	}
}
