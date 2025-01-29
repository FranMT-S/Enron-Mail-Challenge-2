package workers

import (
	"runtime"
	"sync"
)

type IWorkerpool[In any, Out any] interface {
}

type WorkerPoolConfig[Out any] struct {
	ResultQueue chan Out
	ErrorCh     chan error
	Workers     int
}

type WorkerPool[In any, Out any] struct {
	job       Job[In, Out]
	taskQueue chan In
	wg        *sync.WaitGroup
	config    WorkerPoolConfig[Out]
}

func NewWorkerPool[In any, Out any](Job Job[In, Out], TaskQueue chan In, Config WorkerPoolConfig[Out], wg *sync.WaitGroup) *WorkerPool[In, Out] {

	if Config.Workers <= 0 {
		numCores := runtime.NumCPU()
		Config.Workers = numCores
	}

	return &WorkerPool[In, Out]{
		job:       Job,
		config:    Config,
		wg:        wg,
		taskQueue: TaskQueue,
	}
}

func (wp *WorkerPool[In, Out]) Start() {
	for i := 0; i < wp.config.Workers; i++ {
		worker := NewWorker(i, wp.job, wp.taskQueue, wp.config.ResultQueue, wp.config.ErrorCh, wp.wg)
		wp.wg.Add(1)
		go worker.Start()
	}

	wp.wg.Wait()
	close(wp.config.ResultQueue)
}

// func (wp *Worker[In, Out]) SetNumWorkers() {

// 	return wp.
// }
