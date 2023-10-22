package stlbasic

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// Debugger 可debug的
type Debugger interface {
	Debug(prefix uint) string
}

func debugWithPrefix[T any](prefix uint, v T) string {
	if vv, ok := any(v).(Debugger); ok {
		return vv.Debug(prefix)
	} else {
		vv, vt := reflect.ValueOf(v), reflect.TypeOf(v)
		switch vt.Kind() {
		case reflect.Bool:
			if vv.Bool() {
				return "true"
			} else {
				return "false"
			}
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return strconv.FormatInt(vv.Int(), 10)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			return strconv.FormatUint(vv.Uint(), 10)
		case reflect.Uintptr:
			return "0x" + strconv.FormatUint(vv.Uint(), 16)
		case reflect.Float32, reflect.Float64:
			return strconv.FormatFloat(vv.Float(), 'f', -1, 32)
		case reflect.String:
			return "\"" + vv.String() + "\""
		case reflect.UnsafePointer:
			return "0x" + strconv.FormatUint(uint64(uintptr(vv.UnsafePointer())), 16)
		case reflect.Func:
			if vv.IsNil() {
				return "nil"
			}
			return "func(" + "0x" + strconv.FormatUint(uint64(vv.Pointer()), 16) + ")"
		case reflect.Pointer:
			if vv.IsNil() {
				return "nil"
			}
			return "*" + debugWithPrefix(prefix, vv.Elem().Interface())
		case reflect.Array:
			var buf strings.Builder
			buf.WriteString("array{")
			for i := 0; i < vv.Len(); i++ {
				buf.WriteString(debugWithPrefix(prefix, vv.Index(i).Interface()))
				if i < vv.Len()-1 {
					buf.WriteString(", ")
				}
			}
			buf.WriteByte('}')
			return buf.String()
		case reflect.Slice:
			var buf strings.Builder
			buf.WriteString("slice{")
			for i := 0; i < vv.Len(); i++ {
				buf.WriteString(debugWithPrefix(prefix, vv.Index(i).Interface()))
				if i < vv.Len()-1 {
					buf.WriteString(", ")
				}
			}
			buf.WriteByte('}')
			return buf.String()
		case reflect.Map:
			var buf strings.Builder
			buf.WriteString("map{")
			var i int
			for iter := vv.MapRange(); iter.Next(); {
				i++
				buf.WriteString(debugWithPrefix(prefix, iter.Key().Interface()))
				buf.WriteString(": ")
				buf.WriteString(debugWithPrefix(prefix, iter.Value().Interface()))
				if i < vv.Len()-1 {
					buf.WriteString(", ")
				}
			}
			buf.WriteByte('}')
			return buf.String()
		case reflect.Struct:
			var buf strings.Builder
			buf.WriteString(vt.String())
			buf.WriteByte('{')
			for i := 0; i < vt.NumField(); i++ {
				field := vt.Field(i)
				if !field.IsExported() {
					continue
				}
				buf.WriteByte('\n')
				for j := 0; j <= int(prefix); j++ {
					buf.WriteString("  ")
				}
				buf.WriteString(field.Name)
				buf.WriteString(": ")
				buf.WriteString(debugWithPrefix(prefix+1, vv.Field(i).Interface()))
			}
			buf.WriteByte('\n')
			for j := 0; j < int(prefix); j++ {
				buf.WriteString("  ")
			}
			buf.WriteByte('}')
			return buf.String()
		default:
			panic(fmt.Errorf("type `%s` cannot be get debug string", vt))
		}
	}
}

// Debug 获取debug字符串
func Debug[T any](v T, prefix ...uint) string {
	var skip uint
	if len(prefix) > 0 {
		skip = prefix[0]
	}
	return debugWithPrefix[T](skip, v)
}
