package stlerror

import "fmt"

func RecoverTo(ptr *Error) {
	errObj := recover()
	if errObj == nil {
		return
	}

	switch err := errObj.(type) {
	case error:
		*ptr = ErrorWrap(err)
	case fmt.Stringer:
		*ptr = Errorf("%s", err.String())
	default:
		*ptr = Errorf("%v", err)
	}
}
