package shared

import (
	"log"
	"os"
)

func CreateFolderIfNotExist(dir string) {
	err := os.MkdirAll(dir, 0755) // 0755 da permisos de lectura y ejecución a otros
	if err != nil {
		log.Println("Error trying to create the folder "+dir, err)
		return
	}
}
