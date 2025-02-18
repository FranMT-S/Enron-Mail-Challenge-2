package shared

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

// Recorrer directorios de manera concurrente y enviar las rutas de los archivos a través de un canal
func FindFilesAsync(dir string, fileChan chan<- string, wg *sync.WaitGroup) {
	// Defer la llamada a Done() para que se ejecute cuando la función termine
	defer wg.Done()

	// Abrir el directorio
	files, err := os.ReadDir(dir)
	if err != nil {
		fmt.Printf("Error al abrir el directorio %s: %v\n", dir, err)
		return
	}

	// Recorrer los archivos y subdirectorios
	for _, file := range files {
		filePath := filepath.Join(dir, file.Name())
		if file.IsDir() {
			// Si es un subdirectorio, procesarlo en una nueva goroutine
			wg.Add(1)                                 // Incrementar el contador del WaitGroup para la nueva goroutine
			go FindFilesAsync(filePath, fileChan, wg) // Llamada recursiva en goroutine
		} else {
			fileChan <- filePath
		}
	}
}
