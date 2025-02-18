package shared

type semaphore struct {
	ch chan struct{}
}

// Limit 0 is equals to mutext
func NewSemaphore(limit int) *semaphore {
	return &semaphore{
		ch: make(chan struct{}, limit),
	}
}

func (semaphore *semaphore) Signal() {
	<-semaphore.ch
}

func (semaphore *semaphore) Wait() {
	semaphore.ch <- struct{}{}
}
