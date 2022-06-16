package types

// Option 选项
type Option[T any] struct {
	ok    bool // 是否有值
	value T    // 值
}

// Some 有值
func Some[T any](v T) Option[T] {
	return Option[T]{
		ok:    true,
		value: v,
	}
}

// None 无值
func None[T any]() Option[T] {
	return Option[T]{}
}

// IsNone 是否是空值
func (self Option[T]) IsNone() bool {
	return !self.ok
}

// IsSome 是否有值
func (self Option[T]) IsSome() bool {
	return self.ok
}

// Unwrap 解包
func (self Option[T]) Unwrap() T {
	if !self.ok {
		panic("expect a value but there is none")
	}
	return self.value
}

// UnwrapOr 解包否则
func (self Option[T]) UnwrapOr(v T) T {
	if !self.ok {
		return v
	}
	return self.value
}

// UnwrapOrDefault 解包否则默认值
func (self Option[T]) UnwrapOrDefault() T {
	return self.value
}

// Take 获取并赋空
func (self *Option[T]) Take() T {
	if !self.ok {
		panic("expect a value but there is none")
	}
	self.ok = false
	var v2 T
	v := self.value
	self.value = v2
	return v
}
