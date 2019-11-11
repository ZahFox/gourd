package common

import (
	"os"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/zahfox/gourd/pkg/utils"
)

// EnvConfigOpts are options used for the environment configuration
type EnvConfigOpts struct {
	Name   string
	Path   string
	Prefix string
	Write  func() error
}

// EnvConfig environment dependent configurations shared between applications
func EnvConfig(env utils.Env, opts EnvConfigOpts) {
	setupLogging()

	viper.SetConfigName(opts.Name)
	viper.AddConfigPath(opts.Path)
	viper.SetEnvPrefix(opts.Prefix)
	viper.BindEnv("SOCKET")

	switch env {
	case utils.DebugEnv:
		{
			utils.MkdirIfNotExist("/tmp/.gourd")
			viper.SetDefault("SOCKET", "/tmp/.gourd/gourdd-debug.sock")
			break
		}
	case utils.ProdEnv:
		{
			utils.MkdirIfNotExist(opts.Path)
			viper.SetDefault("SOCKET", "/run/gourd/gourdd.sock")
			break
		}
	case utils.DevEnv:
		{
			utils.MkdirIfNotExist(opts.Path)
			viper.SetDefault("SOCKET", "/run/gourd/gourdd.sock")
			break
		}
	}

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			err = opts.Write()
			if err != nil {
				utils.LogFatalf("no configuration file was found for %s", strings.ToLower(opts.Name))
			}

			if err := viper.ReadInConfig(); err != nil {
				if _, ok := err.(viper.ConfigFileNotFoundError); ok {
					utils.LogFatalf("no configuration file was found for %s", strings.ToLower(opts.Name))
				} else {
					utils.LogFatalf("the configuration file for %s is invalid", strings.ToLower(opts.Name))
				}
			}

		} else {
			utils.LogFatalf("the configuration file for %s is invalid", strings.ToLower(opts.Name))
		}
	}
}

// GetSocketPath returns the filesystem path to the command socket
func GetSocketPath() string {
	return viper.GetString("SOCKET")
}

func setupLogging() {
	utils.SetupLogging(func(sol *logrus.Logger, sel *logrus.Logger) {
		sol.SetFormatter(&logrus.TextFormatter{})
		sol.SetOutput(os.Stdout)
		sel.SetFormatter(&logrus.TextFormatter{})
		sel.SetOutput(os.Stderr)
		env := utils.GetEnv()
		switch env {
		case utils.ProdEnv:
			sol.SetLevel(logrus.InfoLevel)
			break
		default:
			sol.SetLevel(logrus.DebugLevel)
		}
	})
}
