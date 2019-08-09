// Package utils is a collection of misc tools shared by each package
// of the gourd module
package utils

import (
	"bufio"
	"encoding/json"
	"io/ioutil"
	"os"
)

// WriteJSON saves a data structure to a path in JSON format
func WriteJSON(path string, data interface{}) error {
	bytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	bytes = append(bytes, byte('\n'))
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0640)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	_, err = writer.Write(bytes)

	if err != nil {
		writer.Reset(writer)
		return err
	}

	err = writer.Flush()
	if err != nil {
		return err
	}

	err = file.Sync()
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
