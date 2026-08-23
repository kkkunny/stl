package stltype

import (
	"reflect"

	stlval "github.com/kkkunny/stl/value"
)

// ImplInterface 类型是否实现了某个接口
func ImplInterface[Type, Interface any]() bool {
	it := reflect.TypeFor[Interface]()
	if it.Kind() != reflect.Interface {
		_, ok := any(stlval.Default[Type]()).(Interface)
		return ok
	}
	return reflect.TypeFor[Type]().Implements(it)
}