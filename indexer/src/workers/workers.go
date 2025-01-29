package workers

import "sync"

type Worker[In any, Out any] struct {
	id            int
	job           Job[In, Out]
	taskQueueCh   <-chan In
	resultQueueCh chan<- Out
	errorCh       chan error
	wg            *sync.WaitGroup
}

func NewWorker[In any, Out any](
	Id int,
	Job Job[In, Out],
	TaskQueue chan In,
	ResultQueue chan Out,
	ErrorCh chan error,
	wg *sync.WaitGroup,
) *Worker[In, Out] {

	return &Worker[In, Out]{
		id:            Id,
		job:           Job,
		taskQueueCh:   TaskQueue,
		resultQueueCh: ResultQueue,
		errorCh:       ErrorCh,
		wg:            wg,
	}
}

func (wo *Worker[In, Out]) Start() {
	defer wo.wg.Done()

	for task := range wo.taskQueueCh {
		result, err := wo.job(task)
		if err != nil && wo.errorCh != nil {
			wo.errorCh <- err
			continue
		}

		if wo.resultQueueCh != nil {
			wo.resultQueueCh <- result
		}
	}
}
