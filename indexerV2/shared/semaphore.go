package shared

type Semaphore struct {
	ch chan struct{}
}

// Limit 0 is equals to mutext
func NewSemaphore(limit int) *Semaphore {
	return &Semaphore{
		ch: make(chan struct{}, limit),
	}
}

func (semaphore *Semaphore) Signal() {
	<-semaphore.ch
}

func (semaphore *Semaphore) Wait() {
	semaphore.ch <- struct{}{}
}
