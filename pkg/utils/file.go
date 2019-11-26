// Package utils is a collection of misc tools shared any gourd application
package utils

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// FileError is used to tell the difference between different kind of file errors
type FileError int

const (
	// FENONE means there is no file error
	FENONE FileError = iota
	// FENOEXIST means a file or directory does not exist
	FENOEXIST
	// FEINSFPERM means the active user has insufficient permissions for the
	// requested action on the file
	FEINSFPERM
)

// Exists returns whether the given file or directory exists
func Exists(path string) (bool, FileError) {
	_, err := os.Stat(path)
	if err == nil {
		return true, FENONE
	}

	if os.IsNotExist(err) {
		return false, FENOEXIST
	}

	if os.IsPermission(err) {
		return true, FEINSFPERM
	}

	return true, FENONE
}

// MkdirIfNotExist is like mkdir -p
func MkdirIfNotExist(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0750)
		if err != nil {
			panic(err)
		}
	}
}

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
