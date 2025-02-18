package workers

import (
	"fmt"
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
	job         Job[In, Out]
	taskQueue   chan In
	wg          sync.WaitGroup
	config      WorkerPoolConfig[Out]
	workersList []*Worker[In, Out]
	Name        string
}

func NewWorkerPool[In any, Out any](Job Job[In, Out], TaskQueue chan In, Config WorkerPoolConfig[Out]) *WorkerPool[In, Out] {

	if Config.Workers <= 0 {
		numCores := runtime.NumCPU()
		Config.Workers = numCores
	}

	return &WorkerPool[In, Out]{
		job:       Job,
		config:    Config,
		taskQueue: TaskQueue,
	}
}

func (wp *WorkerPool[In, Out]) SetName(name string) *WorkerPool[In, Out] {
	wp.Name = name
	return wp
}

func (wp *WorkerPool[In, Out]) Start() {
	for i := 0; i < wp.config.Workers; i++ {
		worker := NewWorker(i, wp.job, wp.taskQueue, wp.config.ResultQueue, wp.config.ErrorCh, &wp.wg)
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
		if wp.config.ResultQueue != nil {
			close(wp.config.ResultQueue)
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
	fmt.Printf("Indexer '%v',Sucess task:%v , failed task:%v\n", wp.Name, sucess, failed)
}
