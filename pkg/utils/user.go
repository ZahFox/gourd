package utils

import (
	"os/user"
)

var currentUser *user.User

func init() {
	var err error
	if currentUser, err = user.Current(); err != nil {
		panic(err)
	}
}

// HomeDir returns the filesystem path to the current user's $HOME directory
func HomeDir() string {
	return currentUser.HomeDir
}

// Username returns the username of the user that owns the running process
func Username() string {
	return currentUser.Username
}
