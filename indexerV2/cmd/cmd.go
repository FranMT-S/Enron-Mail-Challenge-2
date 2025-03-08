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
	var wg, wgBatchMails sync.WaitGroup
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

	semaphore := shared.NewSemaphore(50)
	mailUploader := indexer.NewMailsUploader(config.CFG.DatabaseName, db.ZincDatabase())
	mimeIndexer := indexer.NewMimeIndexer()
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
	).SetName("Indexer Mails")

	uploaderPool := workers.NewWorkerPool(
		mailUploader.Upload,
		mailsBatchQueoeCh,
		nil,
		workers.WorkerPoolConfig{
			Workers: configProgram.workersUploaders,
			ErrorCh: errorHandler.GetErrCh(),
		},
	).SetName("Uploader Pool")

	workerpool.Start()
	go indexer.BatchQueoue(configProgram.batchSize, mailsQueoeCh, mailsBatchQueoeCh, &wgBatchMails)
	uploaderPool.Start()

	workerpool.Wait()
	wgBatchMails.Wait()
	uploaderPool.Wait()
	errorHandler.CloseAndWait()
	workerpool.PrintSuccesfulTask()
	uploaderPool.PrintSuccesfulTask()

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
		newDir := strings.TrimSpace(scanner.Text())

		if strings.ToLower(newDir) == "x" {
			os.Exit(0)
		}

		if shared.ExistDirectory(newDir) {
			return newDir
		}
		shared.PrintlnWithColor(shared.TextYellow, "directory not valid.")
	}
}

// searchs files of mails async and add a new value in the wait group param.
// semaphore limit the number of gourotines finding for new mails
func StartSearchMails(dir string, fileChan chan<- string, wg *sync.WaitGroup, semaphore *shared.Semaphore) {
	wg.Add(1)
	go func() {
		shared.FindFilesAsync(dir, fileChan, wg, semaphore)
		wg.Wait()
		close(fileChan)
	}()
}
