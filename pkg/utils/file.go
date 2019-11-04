// Package utils is a collection of misc tools shared any gourd application
package utils

import (
	"encoding/json"
	"io/ioutil"
)

// WriteJSON saves a data structure to a path in JSON format
func WriteJSON(path string, data interface{}) error {
	bytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	bytes = append(bytes, byte('\n'))
	return ioutil.WriteFile(path, bytes, 0640)
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
