package enum

import (
	"reflect"
	"strconv"
)

/*
1. 数字类型默认值为字段下标

2. 字符串类型默认值为字段名

3. 可以通过enum标签指定字段的值

	New[struct {
		A int                // value: 0
		B uint               // value: 1
		C string             // value: "C"
		D string `enum:"A"`  // value: "A"
		E float32            // value: 4.0
	}]()
*/
func New[T any]() T {
	t := reflect.TypeFor[T]()
	if t.Kind() != reflect.Struct {
		panic("expect a struct type")
	}

	v := reflect.New(t).Elem()
	for i := 0; i < t.NumField(); i++ {
		f, fv := t.Field(i), v.Field(i)
		tag, existTag := f.Tag.Lookup("enum")
		switch f.Type.Kind() {
		case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int:
			if tagV, ok := strconv.ParseInt(tag, 10, 64); ok == nil {
				fv.SetInt(tagV)
			} else {
				fv.SetInt(int64(i))
			}
		case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uint, reflect.Uintptr:
			if tagV, ok := strconv.ParseUint(tag, 10, 64); ok == nil {
				fv.SetUint(tagV)
			} else {
				fv.SetUint(uint64(i))
			}
		case reflect.Float32, reflect.Float64:
			if tagV, ok := strconv.ParseFloat(tag, 64); ok == nil {
				fv.SetFloat(tagV)
			} else {
				fv.SetFloat(float64(i))
			}
		case reflect.String:
			if existTag {
				fv.SetString(tag)
			} else {
				fv.SetString(f.Name)
			}
		}
	}
	return v.Interface().(T)
}
