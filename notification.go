package gentypes

type Notification struct {
	promise *Promise[struct{}]
}

func (n Notification) Chan() <-chan struct{} {
	return n.promise.Result()
}

func (n Notification) Done() {
	n.promise.Apply()
}

func (n Notification) Wait() {
	<-n.Chan()
}

func NewNotification() Notification {
	return Notification{NewPromise(func() struct{} {
		return struct{}{}
	})}
}
