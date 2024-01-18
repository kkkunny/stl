package stlerror

// Ignore 忽略异常
func Ignore(_ error) {}

// IgnoreWith 忽略异常
func IgnoreWith[T any](v T, _ error) T {
	return v
}

// IgnoreWith2 忽略异常
func IgnoreWith2[T, E any](v1 T, v2 E, _ error) (T, E) {
	return v1, v2
}

// IgnoreWith3 忽略异常
func IgnoreWith3[T, E, F any](v1 T, v2 E, v3 F, _ error) (T, E, F) {
	return v1, v2, v3
}
