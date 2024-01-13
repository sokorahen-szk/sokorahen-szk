package cmd

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func ReadConfig(jsonFilePath string, v interface{}) error {
	jsonFile, err := os.Open(jsonFilePath)
	if err != nil {
		return err
	}
	defer jsonFile.Close()

	jsonData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return err
	}

	return json.Unmarshal(jsonData, &v)
}
