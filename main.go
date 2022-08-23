package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

func main() {
	args := os.Args
	if len(os.Args) <= 1 {
		FailWithMessage("You must pass at least one file!")
	}

	for _, fileName := range args[1:] {
		if _, err := os.Stat(fileName); errors.Is(err, os.ErrNotExist) {
			FailedToConvertMessage(fmt.Sprint("File '", fileName, "' does not exist!"))
			continue
		}

		yamlData, err := ioutil.ReadFile(fileName)
		if err != nil {
			FailedToConvertMessage(err.Error())
			continue
		}

		var data any = map[interface{}]interface{}{}
		err = yaml.Unmarshal(yamlData, &data)
		if err != nil {
			FailedToConvertMessage(err.Error())
			continue
		}

		jsonData, err := json.Marshal(data)
		if err != nil {
			FailedToConvertMessage(err.Error())
			continue
		}

		jsonName := fmt.Sprint(strings.TrimSuffix(fileName, filepath.Ext(fileName)), ".json")

		err = ioutil.WriteFile(jsonName, jsonData, 0644)
		if err != nil {
			FailedToConvertMessage(err.Error())
			continue
		}
	}

}

func FailedToConvertMessage(message string) {
	log.Print("Failed to convert file: ", message, "\n")
}

func FailWithMessage(message string) {
	log.Println(message)
	os.Exit(1)
}
