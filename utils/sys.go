package utils

import (
	"bytes"
	"os/exec"
	"strings"
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
