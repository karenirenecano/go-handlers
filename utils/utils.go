package utils

import (
	"fmt"
	"log"
	"os"
)

func GetCWD(file string) (certDir string, err error) {
	path, errorNotFound := os.Getwd()
	if errorNotFound != nil {
		log.Fatal(errorNotFound)
	}
	fileName := path + file
	_, errorMessage := os.Stat(fileName)
	if os.IsNotExist(errorMessage) {
		return "Not existing", fmt.Errorf("file [%s] does not exist", fileName)
	}

	return fileName, nil
}
