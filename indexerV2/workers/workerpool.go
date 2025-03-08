package workers

import (
	"fmt"
	"runtime"
	"sync"
)

type IWorkerpool[In any, Out any] interface {
}

type WorkerPoolConfig struct {
	ErrorCh chan error
	Workers int
}

type WorkerPool[In any, Out any] struct {
	job         Job[In, Out]
	taskQueue   chan In
	ResultQueue chan Out
	wg          sync.WaitGroup
	config      WorkerPoolConfig
	workersList []*Worker[In, Out]
	Name        string
}

/*
	Create a new instance of Workerpool

Params:

  - Job a function that return a value of In and an error

  - TaskQueue chan with data that process by Job

  - ResultQueue optional chan with the result of Job

    Config Config of worked pool:

  - ErrorCh optional chan to receive the errors ​​of Job

  - Workers number of workers to process Job
*/
func NewWorkerPool[In any, Out any](Job Job[In, Out], TaskQueue chan In, ResultQueue chan Out, Config WorkerPoolConfig) *WorkerPool[In, Out] {

	if Config.Workers <= 0 {
		numCores := runtime.NumCPU()
		Config.Workers = numCores
	}

	return &WorkerPool[In, Out]{
		job:         Job,
		config:      Config,
		taskQueue:   TaskQueue,
		ResultQueue: ResultQueue,
	}
}

func (wp *WorkerPool[In, Out]) SetName(name string) *WorkerPool[In, Out] {
	wp.Name = name
	return wp
}

func (wp *WorkerPool[In, Out]) Start() {
	for i := 0; i < wp.config.Workers; i++ {
		worker := NewWorker(i, wp.job, wp.taskQueue, wp.ResultQueue, wp.config.ErrorCh, &wp.wg)
		wp.wg.Add(1)
		wp.workersList = append(wp.workersList, worker)
		go worker.Start()
	}
	wp.CloseResultChannelWhenFinish()
}

func (wp *WorkerPool[In, Out]) Wait() {
	wp.wg.Wait()
}

func (wp *WorkerPool[In, Out]) CloseResultChannelWhenFinish() {
	go func() {
		wp.wg.Wait()
		if wp.ResultQueue != nil {
			close(wp.ResultQueue)
		}
	}()
}

func (wp *WorkerPool[In, Out]) GetTasksuccessful() (sucess, failed int) {
	for _, worker := range wp.workersList {
		sucess += worker.successTask
		failed += worker.failedTask
	}

	return sucess, failed
}

func (wp *WorkerPool[In, Out]) PrintSuccesfulTask() {
	sucess, failed := wp.GetTasksuccessful()
	fmt.Printf("Worker '%v',Sucess task:%v , failed task:%v\n", wp.Name, sucess, failed)
}
