package gentypes

type Optional[T any] struct {
	value    T
	hasValue bool
}

func Some[T any](val T) Optional[T] {
	return Optional[T]{value: val, hasValue: true}
}

func None[T any]() Optional[T] {
	return Optional[T]{hasValue: false}
}

func (o Optional[T]) IsSome() bool {
	return o.hasValue
}

func (o Optional[T]) IsNone() bool {
	return !o.hasValue
}

func (o Optional[T]) Value() T {
	if o.hasValue {
		return o.value
	}
	panic("Optional.Value() called on None")
}

func (o Optional[T]) OrElse(other T) T {
	if o.hasValue {
		return o.value
	}
	return other
}

func OptionalTransform[T1 any, T2 any](o Optional[T1], f func(T1) T2) Optional[T2] {
	if o.hasValue {
		return Some(f(o.value))
	}
	return None[T2]()
}
