package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/FranMT-S/Enron-Mail-Challenge-2/indexer/config"
	db "github.com/FranMT-S/Enron-Mail-Challenge-2/indexer/database"
	apperrors "github.com/FranMT-S/Enron-Mail-Challenge-2/indexer/errors"
	"github.com/FranMT-S/Enron-Mail-Challenge-2/indexer/indexer"
	"github.com/FranMT-S/Enron-Mail-Challenge-2/indexer/models"
	"github.com/FranMT-S/Enron-Mail-Challenge-2/indexer/profiling"
	"github.com/FranMT-S/Enron-Mail-Challenge-2/indexer/shared"
	"github.com/FranMT-S/Enron-Mail-Challenge-2/indexer/workers"
)

var _time = time.Now().Format("010206_030405")

func main() {
	config.LoadConfig()
	var wg sync.WaitGroup

	zincDb := db.ZincDatabase()
	zincDb.CreateIndex(config.CFG.DatabaseName)
	mailsQueoeCh := make(chan *models.Email)
	profiler := profiling.Profiler{}

	errorHandler := apperrors.NewErrorHandler()
	errorHandler.Listen()

	// register time
	startTime := time.Now()

	if config.IS_PROFILER {
		profiler.StartMemAndCPUProfiler()
	}

	if config.IS_PROFILER_TRACER {
		profiler.StartTracerProfiler()
	}

	filesQueoCh := make(chan string)

	wg.Add(1)
	go func() {
		shared.FindFilesAsync("../mails/mailsdir", filesQueoCh, &wg)
		wg.Wait()
		close(filesQueoCh)
	}()

	mimeIndexer := indexer.NewMimeIndexer()
	workerpool := workers.NewWorkerPool(
		mimeIndexer.IndexMail,
		filesQueoCh,
		workers.WorkerPoolConfig[*models.Email]{
			Workers:     20,
			ResultQueue: mailsQueoeCh,
			ErrorCh:     errorHandler.GetErrCh(),
		},
	)

	workerpool.SetName("Indexer Mails")
	workerpool.Start()

	var wgBatch sync.WaitGroup

	wgBatch.Add(1)
	go indexer.PoolMailBatch(mailsQueoeCh, &wgBatch)

	workerpool.Wait()
	errorHandler.CloseAndWait()
	wgBatch.Wait()
	workerpool.PrintSuccesfulTask()

	if config.IS_PROFILER {
		profiler.StopMemAndCPUProfiler()
	}

	if config.IS_PROFILER_TRACER {
		profiler.StopTracerProfiler()
	}

	endTime := time.Now()
	duration := endTime.Sub(startTime)
	seconds := duration.Seconds()
	minutes := duration.Minutes()
	// fmt.Println(len(MailResultQueoCh))
	fmt.Printf("The code ran in %.2f seconds\n", seconds)
	fmt.Printf("The code ran in %.2f minutes\n", minutes)
}
