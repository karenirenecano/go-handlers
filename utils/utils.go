package utils

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

//GetCWD : Get Current Working Directory + file path to be checked
func GetCWD(file string) (filePath string, err error) {
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

//SetGlobalEnvVariables : Reference from .env file the variables to be set os.Setenv
func SetGlobalEnvVariables() {
	envConfig, err := GetCWD("/.env")
	if err != nil {
		log.Fatal(err)
	}
	data, err := ioutil.ReadFile(envConfig)
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	configLoaded := string(data)
	contents := strings.Split(configLoaded, "\n")
	configMap := map[string]string{}
	for k, value := range contents {
		configMap[string(k)] = value
		keyValue := strings.Split(value, "=")
		err := os.Setenv(keyValue[0], keyValue[1])
		if err != nil {
			fmt.Println(err)
		}
	}
}
