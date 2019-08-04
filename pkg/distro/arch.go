package distro

import (
	"github.com/zahfox/gourd/pkg/utils"
)

// archPackageManager is the package manager for Arch Linux
type archPackageManager struct {
	distroType Type
}

func (*archPackageManager) Install(packages ...string) bool {
	args := append([]string{"--noconfirm", "-S"}, packages...)
	return utils.Exec("pacman", args...)
}

func (*archPackageManager) Uninstall(packages ...string) bool {
	args := append([]string{"--noconfirm", "-Rns"}, packages...)
	return utils.Exec("pacman", args...)
}
