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

func HasFunc(f, target, from, to reflect.Type) bool {
	if !f.AssignableTo(ReplaceFuncAnyTo(target, from)) {
		return false
	}
	for i := 0; i < f.NumIn(); i++ {
		inType := f.In(i)
		if from.String() != inType.String() {
			continue
		}
		if !to.AssignableTo(inType) {
			return false
		}
	}
	for i := 0; i < f.NumOut(); i++ {
		outType := f.Out(i)
		if from.String() != outType.String() {
			continue
		}
		if !to.AssignableTo(outType) {
			return false
		}
	}
	return true
}

func HasMethod(t reflect.Type, m reflect.Method, from reflect.Type) bool {
	for i := 0; i < t.NumMethod(); i++ {
		method := t.Method(i)
		if method.Name == m.Name {
			if t.Kind() == reflect.Interface {
				if HasFunc(method.Type, m.Type, from, t) {
					return true
				}
			} else {
				in, out := make([]reflect.Type, method.Type.NumIn()), make([]reflect.Type, method.Type.NumOut())
				for j := 0; j < len(in); j++ {
					in[j] = method.Type.In(j)
				}
				for j := 0; j < len(out); j++ {
					out[j] = method.Type.Out(j)
				}
				if HasFunc(reflect.FuncOf(in[1:], out, method.Type.IsVariadic()), m.Type, from, t) {
					return true
				}
			}
		}
	}
	return false
}

func HasAllMethod(t reflect.Type, from reflect.Type, ms ...reflect.Method) bool {
	for _, m := range ms {
		if !HasMethod(t, m, from) {
			return false
		}
	}
	return true
}
