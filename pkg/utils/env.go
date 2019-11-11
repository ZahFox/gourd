package utils

import (
	"os"
	"strings"
)

// Env is an integer that represents a particular Linux distribution
type Env uint8

const defaultEnv Env = ProdEnv

const (
	// UnknownEnv Environment
	UnknownEnv Env = 0
	// DebugEnv is a debugging Environment
	DebugEnv Env = 1
	// DevEnv is a development Environment
	DevEnv Env = 2
	// ProdEnv is a production Environment
	ProdEnv Env = 3
)

var env Env = UnknownEnv
var socketPath = ""

// GetEnv finds the type of runtime environment that is currently active
func GetEnv() Env {
	if env == UnknownEnv {
		env = processEnv()
	}
	return env
}

// EnvStr returns a human readable form of value for the current environment
func EnvStr() string {
	switch GetEnv() {
	case ProdEnv:
		return "production"
	case DevEnv:
		return "development"
	case DebugEnv:
		return "debug"
	}
	return "unknown"
}

func processEnv() Env {
	currentEnv := os.Getenv("ENV")
	if currentEnv == "" {
		return defaultEnv
	}

	currentEnv = strings.ToLower(currentEnv)
	switch currentEnv {
	case "prod":
	case "production":
		return ProdEnv
	case "dev":
	case "development":
		return DevEnv
	case "debug":
		return DebugEnv
	}

	return defaultEnv
}
