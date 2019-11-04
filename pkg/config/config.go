package config

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
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
//
// This is used instead of an init function because some interactions with gourd do not
// require access to the gourd config. It is the responsibility of any sigificant system
// to manually invoke Load.
func Load() {
	if initialized {
		return
	}

	checkConfigDir()
	loadConfig()
	setupLogging()
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
	err := utils.ReadJSON(path, data)

	if err != nil {
		data = getDefaultConfig()
		utils.WriteJSON(path, data)
	}

	// TODO:
	// 1. Validate the config file
	// 2. Add logic for configs with outdated or invalid versions
	return data
}

func setupLogging() {
	utils.SetupLogging(func(sol *log.Logger, sel *log.Logger) {
		sol = log.New()
		sol.SetFormatter(&log.TextFormatter{})
		sol.SetOutput(os.Stdout)

		sel = log.New()
		sel.SetFormatter(&log.TextFormatter{})
		sel.SetOutput(os.Stderr)

		env := GetEnv()
		switch env {
		case Prod:
			sol.SetLevel(log.InfoLevel)
			break
		default:
			sol.SetLevel(log.DebugLevel)
		}
	})
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
