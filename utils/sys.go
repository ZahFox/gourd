package utils

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"strconv"
	"strings"
	"syscall"
)

// OsInfo represents data about the host machine's operating system
type OsInfo struct {
	KernelName      string `json:"kernel_name"`
	NodeName        string `json:"node_name"`
	KernelRelease   string `json:"kernel_release"`
	Machine         string `json:"machine"`
	OperatingSystem string `json:"operating_system"`
}

// Os returns HostInfo about the host machine.
func Os() (OsInfo, error) {
	var host OsInfo
	var out bytes.Buffer

	cmd := exec.Command("uname", "-snrmo")
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		return host, err
	}

	sysinfo := strings.Split(out.String(), " ")
	host = OsInfo{
		KernelName:      sysinfo[0],
		NodeName:        sysinfo[1],
		KernelRelease:   sysinfo[2],
		Machine:         sysinfo[3],
		OperatingSystem: sysinfo[4],
	}

	return host, err
}

// ProgramExists checks to see if a program is in the current $PATH and is executable
func ProgramExists(name string) (string, error) {
	path, err := exec.LookPath(name)
	if err != nil {
		errMsg := fmt.Sprintf("Could not find the executable: %s. Please check your $PATH.", name)
		return "", errors.New(errMsg)
	}

	info, err := os.Stat(path)
	if err != nil {
		errMsg := fmt.Sprintf("Failed to stat the following file: %s.", path)
		return "", errors.New(errMsg)
	}

	userInfo, err := user.Current()
	if err != nil {
		errMsg := fmt.Sprintf("Failed to find information about the current user.")
		return "", errors.New(errMsg)
	}

	mode := info.Mode()
	ox := mode & 1
	or := mode & 4

	if ox == 1 && or == 4 {
		return path, nil
	}

	ux := mode & 64
	ur := mode & 256
	gx := mode & 8
	gr := mode & 32
	fileUID := info.Sys().(*syscall.Stat_t).Uid
	fileGID := info.Sys().(*syscall.Stat_t).Gid
	userUID, _ := strconv.ParseUint(userInfo.Uid, 10, 32)
	userGID, _ := strconv.ParseUint(userInfo.Gid, 10, 32)

	if uint32(userUID) == fileUID && ux == 64 && ur == 256 {
		return path, nil
	} else if uint32(userGID) == fileGID && gx == 8 && gr == 32 {
		return path, nil
	}

	errMsg := fmt.Sprintf("The current user does not have permission to execute: %s.", path)
	return "", errors.New(errMsg)
}
