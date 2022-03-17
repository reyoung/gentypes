package gentypes

import "emperror.dev/errors"

var ErrChanClosed = errors.NewPlain("chan is closed")

func ChanHead[T any](ch <-chan T) Status[T] {
	head, ok := <-ch
	if !ok {
		return NewErrorStatus[T](ErrChanClosed)
	} else {
		return NewStatus(head, nil)
	}
}
