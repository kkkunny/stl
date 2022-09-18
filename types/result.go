package types

// Result 结果
type Result[T, E any] struct {
	hasValue bool // 是否有值
	value    T    // 值
	error    E    // 异常
}

// Ok 有结果
func Ok[T, E any](v T) Result[T, E] {
	return Result[T, E]{
		hasValue: true,
		value:    v,
	}
}

// Err 异常
func Err[T, E any](e E) Result[T, E] {
	return Result[T, E]{error: e}
}

// IsOk 是否有结果
func (self Result[T, E]) IsOk() bool {
	return self.hasValue
}

// IsErr 是否是异常
func (self Result[T, E]) IsErr() bool {
	return !self.hasValue
}

// Ok 获取结果
func (self Result[T, E]) Ok() Option[T] {
	if !self.hasValue {
		return None[T]()
	}
	return Some(self.value)
}

// Err 获取异常
func (self Result[T, E]) Err() Option[E] {
	if self.hasValue {
		return None[E]()
	}
	return Some(self.error)
}

// Unwrap 解包
func (self Result[T, E]) Unwrap() T {
	if !self.hasValue {
		panic(self.error)
	}
	return self.value
}

// UnwrapOr 解包否则
func (self Result[T, E]) UnwrapOr(v T) T {
	if !self.hasValue {
		return v
	}
	return self.value
}

// UnwrapOrDefault 解包否则默认值
func (self Result[T, E]) UnwrapOrDefault() T {
	return self.value
}
