package gentypes

import "emperror.dev/errors"

type Status[T any] struct {
	value T
	err   error
}

func NewStatus[T any](value T, err error) Status[T] {
	return Status[T]{value: value, err: err}
}

func NewErrorStatus[T any](err error) Status[T] {
	return Status[T]{err: err}
}

func (s Status[T]) Value() T {
	if s.err == nil {
		return s.value
	} else {
		panic(s.err)
	}
}

func (s Status[T]) WrapErr(msg string, args ...interface{}) Status[T] {
	if s.err == nil {
		return s
	} else {
		s.err = errors.WrapIff(s.err, msg, args...)
		return s
	}
}

func (s Status[T]) ValueOr(defaultVal T) T {
	if s.err != nil {
		return defaultVal
	} else {
		return s.value
	}
}

func (s Status[T]) Result() (T, error) {
	return s.value, s.err
}

func RecoverError(err *error) {
	r := recover()
	if r == nil {
		return
	}
	e, ok := r.(error)
	if !ok {
		panic(r)
	}
	*err = e
}

func StatusThen[T1 any, T2 any](t1 Status[T1], transform func(T1) (T2, error)) Status[T2] {
	if t1.err != nil {
		return NewErrorStatus[T2](t1.err)
	} else {
		return NewStatus(transform(t1.value))
	}
}

func StatusFlatten[T any](val Status[Status[T]]) Status[T] {
	if val.err != nil {
		return NewErrorStatus[T](val.err)
	} else {
		return val.value
	}
}
