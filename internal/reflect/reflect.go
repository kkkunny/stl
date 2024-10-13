package reflect

import "reflect"

func ReplaceFuncAnyTo(f, to reflect.Type) reflect.Type {
	anyType := TypeFor[any]()

	in, out := make([]reflect.Type, f.NumIn()), make([]reflect.Type, f.NumOut())
	for i := 0; i < len(in); i++ {
		inType := f.In(i)
		if anyType.AssignableTo(inType) {
			in[i] = to
		} else {
			in[i] = inType
		}
	}
	for i := 0; i < len(out); i++ {
		outType := f.Out(i)
		if anyType.AssignableTo(outType) {
			out[i] = to
		} else {
			out[i] = outType
		}
	}
	return reflect.FuncOf(in, out, f.IsVariadic())
}

func ReplaceMethodAnyTo(m reflect.Method, to reflect.Type) reflect.Method {
	m.Type = ReplaceFuncAnyTo(m.Type, to)
	return m
}

func HasMethod(t reflect.Type, m reflect.Method) bool {
	for i := 0; i < t.NumMethod(); i++ {
		method := t.Method(i)
		if method.Name == m.Name {
			if t.Kind() == reflect.Interface && method.Type.AssignableTo(m.Type) {
				return true
			}
			in, out := make([]reflect.Type, method.Type.NumIn()), make([]reflect.Type, method.Type.NumOut())
			for j := 0; j < len(in); j++ {
				in[j] = method.Type.In(j)
			}
			for j := 0; j < len(out); j++ {
				out[j] = method.Type.Out(j)
			}
			if reflect.FuncOf(in[1:], out, method.Type.IsVariadic()).AssignableTo(m.Type) {
				return true
			}
		}
	}
	return false
}

func HasAllMethod(t reflect.Type, ms ...reflect.Method) bool {
	for _, m := range ms {
		if !HasMethod(t, m) {
			return false
		}
	}
	return true
}
