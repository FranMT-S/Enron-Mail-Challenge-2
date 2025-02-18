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

func main() {
	config.LoadConfig()
	var wg sync.WaitGroup

	zincDb := db.ZincDatabase()
	zincDb.CreateIndex(config.CFG.DatabaseName)
	mailsQueoeCh := make(chan *models.Email)
	mailsBatchQueoeCh := make(chan []*models.Email)
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

	semaphore := shared.NewSemaphore(50)
	wg.Add(1)
	go func() {
		shared.FindFilesAsync("../mails/mailsdir", filesQueoCh, &wg, semaphore)
		wg.Wait()
		close(filesQueoCh)
	}()

	mimeIndexer := indexer.NewMimeIndexer()
	workerpool := workers.NewWorkerPool(
		mimeIndexer.IndexMail,
		filesQueoCh,
		mailsQueoeCh,
		workers.WorkerPoolConfig{
			Workers: 20,
			ErrorCh: errorHandler.GetErrCh(),
		},
	)
	workerpool.SetName("Indexer Mails")
	workerpool.Start()

	var wgBatchMails sync.WaitGroup
	var wgUploader sync.WaitGroup

	go indexer.BatchQueoue(1000, mailsQueoeCh, mailsBatchQueoeCh, &wgBatchMails)
	go indexer.UploaderPool(
		4,
		config.CFG.DatabaseName,
		db.ZincDatabase(),
		mailsBatchQueoeCh,
		errorHandler.GetErrCh(),
		&wgUploader,
	)

	workerpool.Wait()
	wgBatchMails.Wait()
	wgUploader.Wait()
	errorHandler.CloseAndWait()
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
