package shared

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

func FindFilesAsync(dir string, fileChan chan<- string, wg *sync.WaitGroup, semaphore *semaphore) {

	defer wg.Done()

	files, err := os.ReadDir(dir)
	if err != nil {
		fmt.Printf("Error al abrir el directorio %s: %v\n", dir, err)
		return
	}

	if semaphore != nil {
		semaphore.Wait()
	}
	for _, file := range files {
		filePath := filepath.Join(dir, file.Name())
		if file.IsDir() {
			wg.Add(1)
			go FindFilesAsync(filePath, fileChan, wg, semaphore)
		} else {
			fileChan <- filePath
		}
	}
	if semaphore != nil {
		semaphore.Signal()
	}
}
