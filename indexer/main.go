package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
	"sync"
	"time"

	"github.com/FranMT-S/Enron-Mail-Challenge-2/indexer/config"
	db "github.com/FranMT-S/Enron-Mail-Challenge-2/indexer/database"
	apperrors "github.com/FranMT-S/Enron-Mail-Challenge-2/indexer/errors"
	"github.com/FranMT-S/Enron-Mail-Challenge-2/indexer/indexer"
	"github.com/FranMT-S/Enron-Mail-Challenge-2/indexer/models"
	"github.com/FranMT-S/Enron-Mail-Challenge-2/indexer/shared"
	"github.com/FranMT-S/Enron-Mail-Challenge-2/indexer/workers"
)

var _time = time.Now().Format("010206_030405")

func main() {
	config.LoadConfig()
	var wg sync.WaitGroup
	var memProf *os.File = nil
	var cpuProf *os.File = nil
	var err error = nil
	zincDb := db.ZincDatabase()
	zincDb.CreateIndex(config.CFG.DatabaseName)
	mailsQueoeCh := make(chan *models.Email)

	errorHandler := apperrors.NewErrorHandler()
	errorHandler.Listen()

	// register time
	startTime := time.Now()
	// filesQueoCh := shared.FindFiles("./mails/mails-1")

	// cpu profiling
	activateProf := true
	if activateProf {
		// cpuProf, err = os.Create(fmt.Sprintf("profiling/cpu_%v.prof", _time))
		cpuProf, err = os.OpenFile("profiling/cpu.prof", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
		if err != nil {
			log.Println("cannot start  CPU profile")
			log.Println(err)
			// _logs.LogSVG(
			// 	constants_log.FILE_NAME_ERROR_PROFILING,
			// 	constants_log.OPERATION_PROFILING,
			// 	constants_log.ERROR_CPU_PROFILING_CREATE,
			// 	err,
			// )
		}

		defer cpuProf.Close() // error handling omitted for example

		if err := pprof.StartCPUProfile(cpuProf); err != nil {
			log.Println("cannot start  CPU profile")
			log.Println(err)
			// _logs.LogSVG(
			// 	constants_log.FILE_NAME_ERROR_PROFILING,
			// 	constants_log.OPERATION_PROFILING,
			// 	constants_log.ERROR_CPU_PROFILING_START,
			// 	err,
			// )
		}

		defer pprof.StopCPUProfile()
	}

	if activateProf {
		// memProf, err = os.Create(fmt.Sprintf("profiling/mem_%v.prof", _time))
		memProf, err = os.OpenFile("profiling/mem.prof", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
		if err != nil {
			log.Println("cannot create memory profile")
			log.Println(err)
			// _logs.LogSVG(
			// 	constants_log.FILE_NAME_ERROR_PROFILING,
			// 	constants_log.OPERATION_PROFILING,
			// 	constants_log.ERROR_MEM_PROFILING_CREATE,
			// 	err,
			// )
		}

		defer memProf.Close()
		runtime.GC()
	}

	filesQueoCh := make(chan string)
	wg.Add(1)
	go func() {
		go shared.FindFilesAsync("./mails/mailsdir", filesQueoCh, &wg)
		wg.Wait()
		close(filesQueoCh)
	}()

	mimeIndexer := indexer.NewMimeIndexer()
	workerpool := workers.NewWorkerPool(
		mimeIndexer.IndexMail,
		filesQueoCh,
		workers.WorkerPoolConfig[*models.Email]{
			Workers:     1000,
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
	if memProf != nil {
		if err := pprof.WriteHeapProfile(memProf); err != nil {
			log.Println("could not write memory profile")
			log.Println(err)
		}
	}

	endTime := time.Now()
	duration := endTime.Sub(startTime)
	seconds := duration.Seconds()
	minutes := duration.Minutes()
	// fmt.Println(len(MailResultQueoCh))
	fmt.Printf("The code ran in %.2f seconds\n", seconds)
	fmt.Printf("The code ran in %.2f minutes\n", minutes)
}
