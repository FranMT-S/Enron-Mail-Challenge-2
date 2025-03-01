package shared

import (
	"os"
)

func ExistDirectory(path string) bool {
	file, err := os.Stat(path)
	if err != nil {
		return false
	}

	return file.IsDir()
}
