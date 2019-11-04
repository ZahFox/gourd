package distro

import (
	"strings"

	"github.com/zahfox/gourd/pkg/utils"
)

// Type is an integer that represents a particular Linux distribution
type Type uint8

const (
	// Arch Linux
	Arch Type = 1
	// CentOS Linux
	CentOS Type = 2
	// Debian Linux
	Debian Type = 3
	// Fedora Linux
	Fedora Type = 4
	// Unknown Linux
	Unknown Type = 255
)

// PackageManager is used to interact with any distribution's package manager
// from a common interface
type PackageManager interface {
	Install(packages ...string) error
	Uninstall(packages ...string) error
}

// Distro is used to interact with any Linux distribution from a common interface
type Distro interface {
	PackageManager
}

type distroInstance struct {
	PackageManager
	DistroType Type
}

var instance distroInstance

// GetDistro will return the Type of the current Linux distribution
func GetDistro() Distro {
	if instance.DistroType != 0 {
		return instance
	}

	osInfo, err := utils.Os()
	if err != nil {
		utils.LogFatal(err)
	}

	if strings.HasSuffix(osInfo.KernelRelease, "ARCH") {
		pkgMgr := archPackageManager{}
		instance = distroInstance{PackageManager: &pkgMgr, DistroType: Arch}
	}

	return instance
}
