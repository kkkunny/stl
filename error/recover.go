package stlerror

func RecoverTo(ptr *Error) {
	errObj := recover()
	if errObj == nil {
		return
	}

	switch err := errObj.(type) {
	case error:
		*ptr = ErrorWrap(err)
	default:
		*ptr = Errorf("%v", err)
	}
}
