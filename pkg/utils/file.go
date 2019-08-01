// Package utils is a collection of misc tools shared by each package
// of the gourd module
package utils

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
)

// WriteJSON saves a data structure to a path in JSON format
func WriteJSON(path string, data interface{}) error {
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

// ReadJSON decodes JSON formatted data from a path into a data structure
func ReadJSON(path string, data interface{}) error {
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
