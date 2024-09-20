package stlerr

import "github.com/pkg/errors"

// CurrentStackTrace 获取目前的错误栈
func CurrentStackTrace() errors.StackTrace {
	return errors.New("").(StackTracer).StackTrace()[1:]
}
