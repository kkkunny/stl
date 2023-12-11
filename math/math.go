package stlmath

import (
	"reflect"
	"strconv"
	"strings"

	"golang.org/x/exp/constraints"

	stlbasic "github.com/kkkunny/stl/basic"
)

// RoundTo 将m取整到n
func RoundTo[T stlbasic.Number](m, n T) T {
	return ((m + n - 1) / n) * n
}

// FIXME: 错误实现
// NotWithBits 按位取反
func NotWithBits[T constraints.Integer](v T, bits uint64) T {
	var bitStr string
	vv := reflect.ValueOf(v)
	if vv.CanInt() {
		bitStr = strconv.FormatInt(vv.Int(), 2)
	} else {
		bitStr = strconv.FormatUint(vv.Uint(), 2)
	}
	if uint64(len(bitStr)) < bits {
		var buf strings.Builder
		for i := uint64(0); i < bits-uint64(len(bitStr)); i++ {
			buf.WriteByte('0')
		}
		buf.WriteString(bitStr)
		bitStr = buf.String()
	} else if uint64(len(bitStr)) > bits {
		bitStr = bitStr[uint64(len(bitStr))-bits:]
	}

	bitList := []byte(bitStr)
	for i, bit := range bitList {
		if bit == '0' {
			bitList[i] = '1'
		} else {
			bitList[i] = '0'
		}
	}
	bitStr = string(bitList)

	var res T
	if vv.CanInt() {
		v, _ := strconv.ParseInt(bitStr, 2, 64)
		res = T(v)
	} else {
		v, _ := strconv.ParseUint(bitStr, 2, 64)
		res = T(v)
	}
	return res
}
