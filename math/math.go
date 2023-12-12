package stlmath

import (
	"reflect"
	"unsafe"

	"golang.org/x/exp/constraints"

	stlbasic "github.com/kkkunny/stl/basic"
	"github.com/samber/lo"
)

// RoundTo 将m取整到n
func RoundTo[T stlbasic.Number](m, n T) T {
	return ((m + n - 1) / n) * n
}

// GetBitFromSignedInt 获取有符号整数指定位数的值（从右向左数）
// eg. 5:00000101 pos=0 ret=1
func GetBitFromSignedInt[T constraints.Signed](v T, pos uint64) bool {
	if pos < uint64(unsafe.Sizeof(v))*8-1{
		val := v & (1 << pos)
		return val > 0
	}else if pos == uint64(unsafe.Sizeof(v))*8-1 {
		return v < 0
	}else{
		return false
	}
}

// GetBitFromUnsignedInt 获取无符号整数指定位数的值（从右向左数）
// eg. 5:00000101 pos=0 ret=1
func GetBitFromUnsignedInt[T constraints.Unsigned](v T, pos uint64) bool {
	if pos < uint64(unsafe.Sizeof(v))*8{
		val := v & (1 << pos)
		return val > 0
	}else{
		return false
	}
}

// GetBitFromInt 获取整数指定位数的值（从右向左数）
// eg. 5:00000101 pos=0 ret=1
func GetBitFromInt[T constraints.Integer](v T, pos uint64) bool {
	switch reflect.TypeOf(T(0)).Kind() {
	case reflect.Int8:
		return GetBitFromSignedInt[int8](int8(v), pos)
	case reflect.Int16:
		return GetBitFromSignedInt[int16](int16(v), pos)
	case reflect.Int32:
		return GetBitFromSignedInt[int32](int32(v), pos)
	case reflect.Int64:
		return GetBitFromSignedInt[int64](int64(v), pos)
	case reflect.Int:
		return GetBitFromSignedInt[int](int(v), pos)
	case reflect.Uint8:
		return GetBitFromUnsignedInt[uint8](uint8(v), pos)
	case reflect.Uint16:
		return GetBitFromUnsignedInt[uint16](uint16(v), pos)
	case reflect.Uint32:
		return GetBitFromUnsignedInt[uint32](uint32(v), pos)
	case reflect.Uint64:
		return GetBitFromUnsignedInt[uint64](uint64(v), pos)
	case reflect.Uint:
		return GetBitFromUnsignedInt[uint](uint(v), pos)
	default:
		panic("unreachable")
	}
}

// GetBitsFromIntWithLength 获取整数所有位
func GetBitsFromIntWithLength[T constraints.Integer](v T, length uint64) []bool {
	res := make([]bool, length)
	for i := uint64(0); i < length; i++ {
		res[length-i-1] = GetBitFromInt(v, i)
	}
	return res
}

// ReverseBits 获取反码
func ReverseBits(bits []bool) []bool {
	return lo.Map(bits, func(item bool, _ int) bool {
		return !item
	})
}

// ComplementBits 获取补码
func ComplementBits(bits []bool) []bool {
	bits = ReverseBits(bits)
	var add bool
	for i := len(bits) - 1; i >= 0; i-- {
		if i == len(bits)-1 || add {
			add = bits[i] == true
			bits[i] = !bits[i]
		} else {
			break
		}
	}
	return bits
}

// SignedIntFromBits 根据每位格式化有符号整数
func SignedIntFromBits[T constraints.Signed](bits []bool) T {
	neg := bits[0]

	if neg {
		bits = ComplementBits(bits)
	}
	num := T(UnsignedIntFromBits[uint64](bits[1:]))

	if !neg {
		return num
	} else {
		return 0 - num
	}
}

// UnsignedIntFromBits 根据每位格式化无符号整数
func UnsignedIntFromBits[T constraints.Unsigned](bits []bool) T {
	var num T
	for i, bit := range bits {
		if i != 0 {
			num <<= 1
		}
		if bit {
			num |= 1
		}
	}
	return num
}

// IntFromBits 根据每位格式化整数
func IntFromBits[T constraints.Integer](bits []bool) T {
	if reflect.ValueOf(T(0)).CanInt() {
		return T(SignedIntFromBits[int64](bits))
	} else {
		return T(UnsignedIntFromBits[uint64](bits))
	}
}

// NotWithBits 按位取反
func NotWithBits[T constraints.Integer](v T, bits uint64) T {
	bs := GetBitsFromIntWithLength(v, bits)
	bs = ReverseBits(bs)
	return IntFromBits[T](bs)
}
