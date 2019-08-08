package utils

import (
	"log"
	"os/user"
)

var currentUser *user.User

func loadUser() {
	if currentUser == nil {
		currUser, err := user.Current()
		if err != nil {
			log.Fatal(err)
		}
		currentUser = currUser
	}
}

// HomeDir returns the filesystem path to the current user's
// $HOME directory
func HomeDir() string {
	loadUser()
	return currentUser.HomeDir
}
