package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// ParseJSONi18File parse string to map
// filePath is path to file with json struct key:value
// return map string:string
func ParseJSONi18File(filePath string) (map[string]string, error) {

	var i18 map[string]string

	// Open file
	jsonFile, err := os.Open(filePath)
	if err != nil {
		return i18, err
	}

	// Close file when done
	defer jsonFile.Close()

	// Read file contents
	fileByte, _ := ioutil.ReadAll(jsonFile)

	// Set json content in map
	json.Unmarshal(fileByte, &i18)

	return i18, nil
}
