package config

import (
	"io/ioutil"
	"os"

	toml "github.com/pelletier/go-toml"
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

	common.EnvConfig(utils.GetEnv(), common.EnvConfigOpts{
		Name:   "gourdd",
		Path:   "/etc/gourd/",
		Prefix: "GOURDD",
		Write: func() error {
			bytes, err := toml.Marshal(Config{
				Version: CurrentVersion,
			})

			if err != nil {
				return err
			}

			path := "/etc/gourd/gourdd.toml"
			err = ioutil.WriteFile(path, bytes, 0640)
			if err != nil {
				return err
			}

			ids, err := utils.GetGourdID()
			if err != nil {
				return err
			}

			return os.Chown(path, ids.UID, ids.GID)
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
