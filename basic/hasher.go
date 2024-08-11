package stlbasic

import (
	"fmt"
	"hash"
	"hash/fnv"
	"math"
	"reflect"
)

// Hashable 可哈希的
type Hashable interface {
	Hash() uint64
}

// Hash 获取哈希
func Hash[T any](v T) uint64 {
	switch vv := any(v).(type) {
	case Hashable:
		return vv.Hash()
	default:
		h := fnv.New64a()
		reflectHash(reflect.ValueOf(v), h)
		return h.Sum64()
	}
}

func reflectHash(v reflect.Value, h hash.Hash64) {
	switch vt := v.Type(); vt.Kind() {
	case reflect.Bool:
		_, _ = h.Write([]byte{boolToBytes(v.Bool())})
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		_, _ = h.Write(int64ToBytes(v.Int()))
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		_, _ = h.Write(uint64ToBytes(v.Uint()))
	case reflect.Float32, reflect.Float64:
		_, _ = h.Write(float64ToBytes(v.Float()))
	case reflect.Complex64, reflect.Complex128:
		_, _ = h.Write(complex128ToBytes(v.Complex()))
	case reflect.String:
		_, _ = h.Write([]byte(v.String()))
	// case reflect.Slice:
	// 	if v.IsNil() {
	// 		return
	// 	}
	// 	fallthrough
	case reflect.Array:
		for i := 0; i < v.Len(); i++ {
			reflectHash(v.Index(i), h)
		}
	case reflect.Chan, reflect.Func, reflect.UnsafePointer, reflect.Pointer, reflect.Map, reflect.Slice:
		_, _ = h.Write(pointerToBytes(v.Pointer()))
	case reflect.Interface:
		if v.IsNil() {
			return
		}
		reflectHash(v.Elem(), h)
	// case reflect.Map:
	// 	if v.IsNil() {
	// 		return
	// 	}
	// 	keys := v.MapKeys()
	// 	for _, k := range keys {
	// 		reflectHash(k, h)
	// 		reflectHash(v.MapIndex(k), h)
	// 	}
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			reflectHash(v.Field(i), h)
		}
	default:
		panic(fmt.Errorf("type `%s` cannot get hash", vt))
	}
}

func boolToBytes(b bool) byte {
	if b {
		return 1
	}
	return 0
}

func int64ToBytes(i int64) []byte {
	return []byte{
		byte(i >> 56), byte(i >> 48), byte(i >> 40), byte(i >> 32),
		byte(i >> 24), byte(i >> 16), byte(i >> 8), byte(i),
	}
}

func uint64ToBytes(i uint64) []byte {
	return []byte{
		byte(i >> 56), byte(i >> 48), byte(i >> 40), byte(i >> 32),
		byte(i >> 24), byte(i >> 16), byte(i >> 8), byte(i),
	}
}

func float64ToBytes(f float64) []byte {
	return uint64ToBytes(math.Float64bits(f))
}

func complex128ToBytes(c complex128) []byte {
	r := float64ToBytes(real(c))
	i := float64ToBytes(imag(c))
	return append(r, i...)
}

func pointerToBytes(p uintptr) []byte {
	return uint64ToBytes(uint64(p))
}
