//go:build !go1.22

package reflect

import "reflect"

func TypeFor[T any]() reflect.Type {
	var v T
	if t := reflect.TypeOf(v); t != nil {
		return t // optimize for T being a non-interface kind
	}
	return reflect.TypeOf((*T)(nil)).Elem() // only for an interface kind
}
