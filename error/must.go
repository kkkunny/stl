package stlerror

// Must 处理异常
func Must(err error) {
	if err != nil {
		panic(err)
	}
}

// MustWith 处理异常
func MustWith[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}

// MustWith2 处理异常
func MustWith2[T, E any](v1 T, v2 E, err error) (T, E) {
	if err != nil {
		panic(err)
	}
	return v1, v2
}

// MustWith3 处理异常
func MustWith3[T, E, F any](v1 T, v2 E, v3 F, err error) (T, E, F) {
	if err != nil {
		panic(err)
	}
	return v1, v2, v3
}
