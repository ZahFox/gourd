package distro

import (
	"github.com/zahfox/gourd/pkg/utils"
)

// archPackageManager is the package manager for Arch Linux
type archPackageManager struct {
	distroType Type
}

func (*archPackageManager) Install(packages ...string) error {
	args := append([]string{"--noconfirm", "-S"}, packages...)
	return utils.Exec("pacman", args...)
}

func (*archPackageManager) Uninstall(packages ...string) error {
	args := append([]string{"--noconfirm", "-Rns"}, packages...)
	return utils.Exec("pacman", args...)
}
