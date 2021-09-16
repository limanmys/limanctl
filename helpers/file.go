package helpers

import (
	"os"
)

func GetFileContents(filepath string) (string, error) {
	file, err := os.ReadFile(filepath)
	return string(file), err
}
