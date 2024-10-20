package stltype

import stlval "github.com/kkkunny/stl/value"

// ImplInterface 类型是否实现了某个接口
func ImplInterface[Type, Interface any]() bool {
	_, ok := any(stlval.Default[Type]()).(Interface)
	return ok
}
