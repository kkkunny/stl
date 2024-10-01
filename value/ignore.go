package stlval

// Ignore 忽略返回值
func Ignore[T any](_ T) {}

// IgnoreWith 忽略返回值
func IgnoreWith[T, E any](v T, _ E) T {
	return v
}

// IgnoreWith2 忽略返回值
func IgnoreWith2[T, E, F any](v1 T, v2 E, _ F) (T, E) {
	return v1, v2
}

// IgnoreWith3 忽略返回值
func IgnoreWith3[T, E, F, G any](v1 T, v2 E, v3 F, _ G) (T, E, F) {
	return v1, v2, v3
}
