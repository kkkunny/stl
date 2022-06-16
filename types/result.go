package types

// Result 结果
type Result[T any] struct {
	error error // 异常
	value T     // 值
}

// Ok 有结果
func Ok[T any](v T) Result[T] {
	return Result[T]{value: v}
}

// Err 异常
func Err[T any](e error) Result[T] {
	if e == nil {
		panic("expect a error but there is nil")
	}
	return Result[T]{error: e}
}

// IsOk 是否有结果
func (self Result[T]) IsOk() bool {
	return self.error == nil
}

// IsErr 是否是异常
func (self Result[T]) IsErr() bool {
	return self.error != nil
}

// Ok 获取结果
func (self Result[T]) Ok() Option[T] {
	if self.error != nil {
		return None[T]()
	}
	return Some(self.value)
}

// Err 获取异常
func (self Result[T]) Err() Option[error] {
	if self.error == nil {
		return None[error]()
	}
	return Some(self.error)
}

// Unwrap 解包
func (self Result[T]) Unwrap() T {
	if self.error != nil {
		panic(self.error)
	}
	return self.value
}

// UnwrapOr 解包否则
func (self Result[T]) UnwrapOr(v T) T {
	if self.error != nil {
		return v
	}
	return self.value
}

// UnwrapOrDefault 解包否则默认值
func (self Result[T]) UnwrapOrDefault() T {
	return self.value
}
