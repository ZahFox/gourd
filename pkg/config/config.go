package config

import (
	"fmt"
	"os"

	"github.com/zahfox/gourd/pkg/utils"
)

const (
	// CurrentVersion identifies the current version of the gourd config
	CurrentVersion = "1.0.0"
)

// Config is a collection of all the necessary configuration
// parameters that gourd uses during runtime
type Config struct {
	Version string `json:"version"`
}

var initialized = false

// Load takes care of all the config package's initialization
func Load() {
	if initialized {
		return
	}

	checkConfigDir()
	loadConfig()
}

func checkConfigDir() {
	path := getConfigDirPath()
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, os.FileMode(0750))
	}
}

func loadConfig() *Config {
	path := getConfigPath()
	var data *Config
	err := utils.ReadJSON(path, &data)

	if err != nil {
		data = getDefaultConfig()
		utils.WriteJSON(path, data)
	}

	return data
}

func getDefaultConfig() *Config {
	return &Config{
		Version: CurrentVersion,
	}
}

func getConfigDirPath() string {
	return fmt.Sprintf("%s/.config/gourd", utils.HomeDir())
}

func getConfigPath() string {
	return fmt.Sprintf("%s/.config/gourd/gourd.json", utils.HomeDir())
}