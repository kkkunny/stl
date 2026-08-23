package json

import (
	"encoding"
	"encoding/json"
	"fmt"
	"io"
	"iter"
	"reflect"
	"strconv"

	"github.com/kkkunny/stl/container/tuple"
)

// MarshalKey 序列化map的键，JSON对象键必须是字符串
func MarshalKey(k any) ([]byte, error) {
	if s, ok := k.(string); ok {
		return json.Marshal(s)
	}
	return json.Marshal(fmt.Sprintf("%v", k))
}

// parseKey 将JSON对象键（恒为string）转换为目标键类型
func parseKey[K any](token any) (K, error) {
	var key K
	s, ok := token.(string)
	if !ok {
		return key, fmt.Errorf("expected JSON object key as string, got `%T`", token)
	}
	kt := reflect.TypeFor[K]()
	convert := func(v reflect.Value) (K, error) {
		if !v.Type().ConvertibleTo(kt) {
			return key, fmt.Errorf("cannot convert JSON object key `%s` to type `%s`", s, kt)
		}
		return v.Convert(kt).Interface().(K), nil
	}
	switch kt.Kind() {
	case reflect.String:
		return convert(reflect.ValueOf(s))
	case reflect.Bool:
		b, err := strconv.ParseBool(s)
		if err != nil {
			return key, fmt.Errorf("cannot convert JSON object key `%s` to type `%s`", s, kt)
		}
		return convert(reflect.ValueOf(b))
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		n, err := strconv.ParseInt(s, 10, kt.Bits())
		if err != nil {
			return key, fmt.Errorf("cannot convert JSON object key `%s` to type `%s`", s, kt)
		}
		return convert(reflect.ValueOf(n))
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		n, err := strconv.ParseUint(s, 10, kt.Bits())
		if err != nil {
			return key, fmt.Errorf("cannot convert JSON object key `%s` to type `%s`", s, kt)
		}
		return convert(reflect.ValueOf(n))
	case reflect.Float32, reflect.Float64:
		f, err := strconv.ParseFloat(s, kt.Bits())
		if err != nil {
			return key, fmt.Errorf("cannot convert JSON object key `%s` to type `%s`", s, kt)
		}
		return convert(reflect.ValueOf(f))
	default:
		if u, ok := any(&key).(encoding.TextUnmarshaler); ok {
			if err := u.UnmarshalText([]byte(s)); err != nil {
				return key, fmt.Errorf("cannot convert JSON object key `%s` to type `%s`", s, kt)
			}
			return key, nil
		}
		return key, fmt.Errorf("cannot convert JSON object key `%s` to type `%s`", s, kt)
	}
}

func UnmarshalToMapObj[K, V any](r io.Reader) iter.Seq2[tuple.Tuple2[K, V], error] {
	dec := json.NewDecoder(r)

	return func(yield func(tuple.Tuple2[K, V], error) bool) {
		t, err := dec.Token()
		if err != nil {
			yield(tuple.Tuple2[K, V]{}, err)
			return
		}
		if delim, ok := t.(json.Delim); !ok || delim != '{' {
			yield(tuple.Tuple2[K, V]{}, fmt.Errorf("expected JSON object start with '{'"))
			return
		}

		for dec.More() {
			keyToken, err := dec.Token()
			if err != nil {
				yield(tuple.Tuple2[K, V]{}, err)
				return
			}
			key, err := parseKey[K](keyToken)
			if err != nil {
				yield(tuple.Tuple2[K, V]{}, err)
				return
			}

			var val V
			if err = dec.Decode(&val); err != nil {
				yield(tuple.Tuple2[K, V]{}, err)
				return
			}

			if !yield(tuple.Pack2[K, V](key, val), nil) {
				return
			}
		}

		t, err = dec.Token()
		if err != nil {
			yield(tuple.Tuple2[K, V]{}, err)
			return
		}
		if delim, ok := t.(json.Delim); !ok || delim != '}' {
			yield(tuple.Tuple2[K, V]{}, fmt.Errorf("expected JSON object end with '}'"))
			return
		}
	}
}
