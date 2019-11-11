package config

import (
	"fmt"

	"github.com/zahfox/gourd/internal/common"
	"github.com/zahfox/gourd/pkg/utils"
)

const (
	// CurrentVersion identifies the current version of the gourd config
	CurrentVersion = "1.0.0"
)

var ready = false

// Config is a collection of all the necessary configuration
// parameters that gourd uses during runtime
type Config struct {
	Version string `json:"version"`
}

func begin() {
	if ready {
		return
	}

	dir := configDir()
	common.EnvConfig(utils.GetEnv(), common.EnvConfigOpts{
		Name:   "gsh",
		Path:   dir,
		Prefix: "GSH",
		Write: func() error {
			return utils.WriteJSON(fmt.Sprintf("%s/gsh.json", dir), Config{
				Version: CurrentVersion,
			})
		},
	})

	ready = true
}

// GetSocketPath returns the filesystem path to the command socket
func GetSocketPath() string {
	if !ready {
		begin()
	}

	return common.GetSocketPath()
}

func configDir() string {
	return fmt.Sprintf("%s/.config/gourd", utils.HomeDir())
}
