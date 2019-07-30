// Package utils is a collection of misc tools shared by each package
// of the gourd module
package utils

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
)

// Write a data structure to a path as JSON
func Write(path string, data interface{}) error {
	var buffer bytes.Buffer
	bytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	buffer.Write(bytes)
	buffer.WriteString("\n")

	writeErr := ioutil.WriteFile(path, buffer.Bytes(), 0600)
	if writeErr != nil {
		return writeErr
	}

	return nil
}

// Read JSON from a path into a data structure
func Read(path string, data interface{}) error {
	bytes, readErr := ioutil.ReadFile(path)
	if readErr != nil {
		return readErr
	}

	err := json.Unmarshal(bytes, &data)
	if err != nil {
		return err
	}

	return nil
}
