package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

var i18Map map[string]string

// ParseJSONi18File parse string to map
// filePath is path to file with json struct key:value
// return map string:string
func parseJSONi18File(filePath string) (map[string]string, error) {

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

// i18 check if key exist & return value
// if not exist return key
func i18(key string) string {

	if key == "" {
		return key
	}

	if val, ok := i18Map[key]; ok {
		return val
	}

	return key

}
