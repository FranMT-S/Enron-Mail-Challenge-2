package apperrors

import (
	"encoding/csv"
	"fmt"
	"os"
	"sync"
	"time"
)

type ErrorHandler struct {
	errCh chan error
	wg    sync.WaitGroup
}

func NewErrorHandler() *ErrorHandler {
	return &ErrorHandler{
		errCh: make(chan error),
	}
}

// Listen error in the channel, must be use closeAndWait must be used to consume errors.
func (errHandler *ErrorHandler) Listen() {
	errHandler.wg.Add(1)
	go func() {
		defer errHandler.wg.Done()
		i := 0
		for err := range errHandler.errCh {
			i++
			LogErrorToCSV(err)
			fmt.Println(fmt.Sprint(i)+"- Error index Mail:", err.Error())
		}
	}()
}

func (errHandler *ErrorHandler) CloseAndWait() {
	close(errHandler.errCh)
	errHandler.wg.Wait()
}

func (errHandler *ErrorHandler) GetErrCh() chan error {
	return errHandler.errCh
}

func (errHandler *ErrorHandler) Submit(err error) {
	errHandler.errCh <- err
}

// logErrorToCSV logs the error and file path to a CSV file.
func LogErrorToCSV(err error) {
	f, fileErr := os.OpenFile("logs/error_log.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if fileErr != nil {
		fmt.Printf("Error opening error log CSV file: %v\n", fileErr)
		return
	}
	defer f.Close() //nolint:errcheck

	writer := csv.NewWriter(f)
	defer writer.Flush()

	fileInfo, errStat := f.Stat()
	if errStat != nil {
		fmt.Printf("Error checking status of error log file: %v\n", err)
		return
	}

	if fileInfo.Size() == 0 {
		err := writer.Write([]string{"timestamp", "errorMessage"})
		if err != nil {
			fmt.Printf("Error writing header to error log CSV file: %v\n", err)
			return
		}
	}

	logTime := time.Now().Format(time.ANSIC)
	record := []string{
		logTime,
		err.Error(),
	}

	if err := writer.Write(record); err != nil {
		fmt.Printf("Error writing error record to CSV log file: %v\n", err)
	}
}
