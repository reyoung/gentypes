package gentypes

type Promise[T any] struct {
	fn         func() T
	resultChan chan T
}

func NewPromise[T any](fn func() T) *Promise[T] {
	if fn == nil {
		panic("NewPromise: fn is nil")
	}
	return &Promise[T]{
		fn:         fn,
		resultChan: make(chan T, 1),
	}
}

func (p *Promise[T]) Result() <-chan T {
	return p.resultChan
}

func (p *Promise[T]) Apply() {
	if p.fn == nil {
		panic("Promise.Apply() should call once")
	}
	p.resultChan <- p.fn()
	p.fn = nil
	close(p.resultChan)
}
