package stlbasic

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"unsafe"
)

// Debugger 可debug的
type Debugger interface {
	Debug(prefix uint) string
}

func debugWithPrefix[T any](prefix uint, v T) string {
	if vv, ok := any(v).(Debugger); ok {
		return vv.Debug(prefix)
	} else {
		vtype := reflect.TypeOf(v)
		switch vtype.Kind() {
		case reflect.Bool:
			if any(v).(bool) {
				return "true"
			} else {
				return "false"
			}
		case reflect.Int:
			return strconv.FormatInt(int64(any(v).(int)), 10)
		case reflect.Int8:
			return strconv.FormatInt(int64(any(v).(int8)), 10)
		case reflect.Int16:
			return strconv.FormatInt(int64(any(v).(int16)), 10)
		case reflect.Int32:
			return strconv.FormatInt(int64(any(v).(int32)), 10)
		case reflect.Int64:
			return strconv.FormatInt(any(v).(int64), 10)
		case reflect.Uint:
			return strconv.FormatUint(uint64(any(v).(uint)), 10)
		case reflect.Uint8:
			return strconv.FormatUint(uint64(any(v).(uint8)), 10)
		case reflect.Uint16:
			return strconv.FormatUint(uint64(any(v).(uint16)), 10)
		case reflect.Uint32:
			return strconv.FormatUint(uint64(any(v).(uint32)), 10)
		case reflect.Uint64:
			return strconv.FormatUint(any(v).(uint64), 10)
		case reflect.Uintptr:
			return "0x" + strconv.FormatUint(uint64(any(v).(uintptr)), 16)
		case reflect.Float32:
			return strconv.FormatFloat(float64(any(v).(float32)), 'f', -1, 32)
		case reflect.Float64:
			return strconv.FormatFloat(any(v).(float64), 'f', -1, 64)
		case reflect.String:
			return "\"" + any(v).(string) + "\""
		case reflect.UnsafePointer:
			return "0x" + strconv.FormatUint(uint64(uintptr(*(*unsafe.Pointer)(unsafe.Pointer(&v)))), 16)
		case reflect.Func:
			vv := reflect.ValueOf(v)
			return "func(" + "0x" + strconv.FormatUint(uint64(vv.Pointer()), 16) + ")"
		case reflect.Pointer:
			vv := reflect.ValueOf(v)
			if vv.IsNil(){
				return "nil"
			}
			return "*" + debugWithPrefix(prefix, vv.Elem().Interface())
		case reflect.Array:
			vv := reflect.ValueOf(v)
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
			vv := reflect.ValueOf(v)
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
			vv := reflect.ValueOf(v)
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
			vv := reflect.ValueOf(v)
			var buf strings.Builder
			buf.WriteString(vtype.String())
			buf.WriteByte('{')
			for i:=0; i<vtype.NumField(); i++ {
				field := vtype.Field(i)
				buf.WriteByte('\n')
				for j:=0; j<=int(prefix); j++{
					buf.WriteString("  ")
				}
				buf.WriteString(field.Name)
				buf.WriteString(": ")
				buf.WriteString(debugWithPrefix(prefix+1, vv.Field(i).Interface()))
			}
			buf.WriteByte('\n')
			for j:=0; j<int(prefix); j++{
				buf.WriteString("  ")
			}
			buf.WriteByte('}')
			return buf.String()
		default:
			panic(fmt.Errorf("type `%s` cannot be get debug string", vtype))
		}
	}
}

// Debug 获取debug字符串
func Debug[T any](v T, prefix ...uint) string {
	var skip uint
	if len(prefix) > 0{
		skip = prefix[0]
	}
	return debugWithPrefix[T](skip, v)
}
