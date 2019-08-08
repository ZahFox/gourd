// Package utils is a collection of misc tools shared by each package
// of the gourd module
package utils

import (
	"bufio"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"os"
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

	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0660)
	if err != nil {
		return err
	}

	writer := bufio.NewWriter(file)
	_, err = writer.Write(buffer.Bytes())

	if err != nil {
		writer.Reset(writer)
		return err
	}

	err = writer.Flush()
	return err
}

// ReadJSON decodes JSON formatted data from a path into a data structure
func ReadJSON(path string, data interface{}) error {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	err = json.Unmarshal(bytes, data)
	if err != nil {
		return err
	}

	return nil
}
