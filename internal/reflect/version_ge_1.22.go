//go:build go1.22

package reflect

import "reflect"

func TypeFor[T any]() reflect.Type {
	return reflect.TypeFor[T]()
}
