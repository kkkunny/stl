package lazy

import (
	"reflect"
	"sync"

	stlcmp "github.com/kkkunny/stl/cmp"
	stlslices "github.com/kkkunny/stl/container/slices"
	stlsync "github.com/kkkunny/stl/sync"
)

// LatestFunc 只保留最新入参的函数执行结果
func LatestFunc[F any](fn F) F {
	f := reflect.ValueOf(fn)
	ft := f.Type()
	if ft.Kind() != reflect.Func {
		panic("expect a function")
	}

	inEqFns := make([]func(l any, r any) bool, ft.NumIn())
	for i := range ft.NumIn() {
		inT := ft.In(i)
		eqFn, ok := stlcmp.GetEqualFuncFromReflect(inT, true)
		if !ok {
			panic("expect a comparable type as function's input")
		}
		inEqFns[i] = eqFn
	}

	var lock stlsync.RWLocker = new(sync.RWMutex)
	var inCaches []any
	var outCaches []reflect.Value

	return reflect.MakeFunc(ft, func(args []reflect.Value) (results []reflect.Value) {
		newIns := stlslices.Map(args, func(i int, arg reflect.Value) any {
			return arg.Interface()
		})
		oldIns, oldOuts := func() ([]any, []reflect.Value) {
			lock.RLock()
			defer lock.RUnlock()
			return inCaches, outCaches
		}()
		eq := len(oldIns) != 0 && stlslices.All(inEqFns, func(i int, eqFn func(l any, r any) bool) bool {
			return eqFn(oldIns[i], newIns[i])
		})
		if eq {
			return oldOuts
		}

		lock.Lock()
		defer lock.Unlock()
		eq = len(inCaches) != 0 && stlslices.All(inEqFns, func(i int, eqFn func(l any, r any) bool) bool {
			return eqFn(inCaches[i], newIns[i])
		})
		if eq {
			return oldOuts
		}

		outs := f.Call(args)
		inCaches, outCaches = newIns, outs
		return outs
	}).Interface().(F)
}
