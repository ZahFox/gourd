package utils

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"strconv"
	"strings"
	"syscall"

	"github.com/pkg/errors"
)

// OsInfo represents data about the host machine's operating system
type OsInfo struct {
	KernelName      string `json:"kernel_name"`
	NodeName        string `json:"node_name"`
	KernelRelease   string `json:"kernel_release"`
	Machine         string `json:"machine"`
	OperatingSystem string `json:"operating_system"`
}

var osInfo OsInfo

// Os returns OsInfo about the host machine.
func Os() (OsInfo, error) {
	if osInfo.KernelName != "" {
		return osInfo, nil
	}

	var host OsInfo
	var out bytes.Buffer

	cmd := exec.Command("uname", "-snrmo")
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		return host, err
	}

	sysinfo := strings.Split(out.String(), " ")
	osInfo = OsInfo{
		KernelName:      sysinfo[0],
		NodeName:        sysinfo[1],
		KernelRelease:   sysinfo[2],
		Machine:         sysinfo[3],
		OperatingSystem: sysinfo[4],
	}

	return osInfo, err
}

// Exec provides a thin wrapper over exec.Command
func Exec(name string, arg ...string) error {
	cmd := exec.Command(name, arg...)
	cmd.Stderr = os.Stdout
	cmd.Stdout = os.Stderr
	return cmd.Run()
}

// UserCanExec checks to see if a file can be executed by the current $USER
func UserCanExec(path string) (bool, error) {
	info, err := os.Stat(path)
	if err != nil {
		return false, errors.Wrapf(err, "Failed to stat the following file: %s", path)
	}

	userInfo, err := user.Current()
	if err != nil {
		return false, errors.Wrapf(err, "Failed to find information about the current user")
	}

	mode := info.Mode()
	ox := mode & 1
	or := mode & 4

	if ox == 1 && or == 4 {
		return true, nil
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
		return true, nil
	} else if uint32(userGID) == fileGID && gx == 8 && gr == 32 {
		return true, nil
	}

	return false, fmt.Errorf("The current user does not have permission to execute: %s.", path)
}
