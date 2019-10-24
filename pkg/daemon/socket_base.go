// +build !systemd

package daemon

import (
	"fmt"
)

// MakeSocket creates the Unix domain socket used for sending commands to gourdd
func MakeSocket() interface{} {
	fmt.Println("NO SYSTEMD!!!")
	return nil
}
