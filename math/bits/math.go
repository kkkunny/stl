package stlbits

import (
	"reflect"

	"golang.org/x/exp/constraints"
)

// Integer 根据每位格式化整数
func Integer[T constraints.Integer](bits Bits) T {
	if reflect.ValueOf(T(0)).CanInt() {
		return T(bits.SignedInteger())
	} else {
		return T(bits.UnsignedInteger())
	}
}

// NotWithLength 按位取反（指定位长）
func NotWithLength[T constraints.Integer](v T, length uint64) T {
	return Integer[T](NewFromIntegerWithLength(v, length).Reverse())
}
