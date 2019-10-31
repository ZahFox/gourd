package config

import (
	"log"
	"os"
	"strings"
)

// Env is an integer that represents a particular Linux distribution
type Env uint8

const defaultEnv Env = Prod

const (
	// Unknown Environment
	Unknown Env = 0
	// Debugging Environment
	Debug Env = 1
	// Development Environment
	Dev Env = 2
	// Production Environment
	Prod Env = 3
)

var env Env = Unknown

// Determine the type of runtime environment that is currently active
func GetEnv() Env {
	if env == Unknown {
		env = processEnv()
		log.Printf("Running in a %s environment", envToString())
	}
	return env
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
		return Prod
	case "dev":
	case "development":
		return Dev
	case "debug":
		return Debug
	}

	return defaultEnv
}

func envToString() string {
	switch env {
	case Prod:
		return "production"
	case Dev:
		return "development"
	case Debug:
		return "debug"
	}
	return "unknown"
}