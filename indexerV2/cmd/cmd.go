package cmd

import (
	"bufio"
	"flag"
	"os"
	"strings"
	"sync"

	"github.com/FranMT-S/Enron-Mail-Challenge-2/indexer/config"
	db "github.com/FranMT-S/Enron-Mail-Challenge-2/indexer/database"
	apperrors "github.com/FranMT-S/Enron-Mail-Challenge-2/indexer/errors"
	"github.com/FranMT-S/Enron-Mail-Challenge-2/indexer/indexer"
	"github.com/FranMT-S/Enron-Mail-Challenge-2/indexer/models"
	"github.com/FranMT-S/Enron-Mail-Challenge-2/indexer/profiling"
	"github.com/FranMT-S/Enron-Mail-Challenge-2/indexer/shared"
	"github.com/FranMT-S/Enron-Mail-Challenge-2/indexer/workers"
)

var scanner *bufio.Scanner = bufio.NewScanner(os.Stdin)

func Start() {
	var wg, wgBatchMails, wgUploader sync.WaitGroup
	var dir string
	configProgram := ConfigProgram{}
	profiler := profiling.Profiler{}

	defineFlag(&configProgram)
	validateFields(&configProgram)
	config.LoadEnviromentConfig()

	args := flag.Args()
	if len(args) > 0 {
		dir = args[0]
	}

	if !shared.ExistDirectory(dir) {
		initialMsg := ""
		if dir == "" {
			initialMsg = "Directory not provided"
		} else {
			initialMsg = "directory not valid"
		}

		shared.PrintlnWithColor(shared.TextYellow, initialMsg)
		dir = CheckDirUntilExist()
	}

	zincDb := db.ZincDatabase()
	zincDb.CreateIndex(config.CFG.DatabaseName)

	errorHandler := apperrors.NewErrorHandler()
	errorHandler.Listen()

	if configProgram.isProf {
		profiler.StartMemAndCPUProfiler()
	}

	if configProgram.isTracer {
		profiler.StartTracerProfiler()
	}

	mimeIndexer := indexer.NewMimeIndexer()
	semaphore := shared.NewSemaphore(50)

	mailsQueoeCh := make(chan *models.Email)
	mailsBatchQueoeCh := make(chan []*models.Email)
	filesQueoCh := make(chan string)

	StartSearchMails(dir, filesQueoCh, &wg, semaphore)

	workerpool := workers.NewWorkerPool(
		mimeIndexer.IndexMail,
		filesQueoCh,
		mailsQueoeCh,
		workers.WorkerPoolConfig{
			Workers: configProgram.workersMails,
			ErrorCh: errorHandler.GetErrCh(),
		},
	)

	workerpool.SetName("Indexer Mails")
	workerpool.Start()

	go indexer.BatchQueoue(configProgram.batchSize, mailsQueoeCh, mailsBatchQueoeCh, &wgBatchMails)
	go indexer.UploaderPool(configProgram.workersUploaders, config.CFG.DatabaseName, db.ZincDatabase(), mailsBatchQueoeCh, errorHandler.GetErrCh(), &wgUploader)

	workerpool.Wait()
	wgBatchMails.Wait()
	wgUploader.Wait()
	errorHandler.CloseAndWait()
	workerpool.PrintSuccesfulTask()

	if configProgram.isProf {
		profiler.StopMemAndCPUProfiler()
	}

	if configProgram.isTracer {
		profiler.StopTracerProfiler()
	}

}

func CheckDirUntilExist() string {
	shared.PrintfWithColor(shared.TextYellow, "send 'x' if you want exit to the program\n")
	for {
		shared.PrintfWithColor(shared.TextGreen, "Enter a directory valid: ")
		scanner.Scan()
		newDir := scanner.Text()

		if strings.ToLower(newDir) == "x" {
			os.Exit(0)
		}

		if shared.ExistDirectory(newDir) {
			return newDir
		}
		shared.PrintlnWithColor(shared.TextYellow, "directory not valid.")
	}
}

func StartSearchMails(dir string, fileChan chan<- string, wg *sync.WaitGroup, semaphore *shared.Semaphore) {
	wg.Add(1)
	go func() {
		shared.FindFilesAsync(dir, fileChan, wg, semaphore)
		wg.Wait()
		close(fileChan)
	}()
}
