package utils

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/zahfox/gourd/pkg/config"
)

var sol *log.Logger
var sel *log.Logger

func init() {
	sol = log.New()
	sol.SetFormatter(&log.TextFormatter{})
	sol.SetOutput(os.Stdout)

	sel = log.New()
	sel.SetFormatter(&log.TextFormatter{})
	sel.SetOutput(os.Stderr)

	env := config.GetEnv()
	switch env {
	case config.Prod:
		sol.SetLevel(log.InfoLevel)
		break
	default:
		sol.SetLevel(log.DebugLevel)
	}

}

// LogDebug logs messages at the debug level
func LogDebug(args ...interface{}) {
	sol.Debug(args...)
}

// LogDebugf logs messages at the debug level
func LogDebugf(msg string, args ...interface{}) {
	sol.Debugf(msg, args...)
}

// LogError logs messages at the error level
func LogError(args ...interface{}) {
	sel.Error(args...)
}

// LogErrorf logs messages at the error level
func LogErrorf(msg string, args ...interface{}) {
	sel.Errorf(msg, args...)
}

// LogError logs messages at the error level
func LogFatal(args ...interface{}) {
	sel.Fatal(args...)
}

// LogErrorf logs messages at the error level
func LogFatalf(msg string, args ...interface{}) {
	sel.Fatalf(msg, args...)
}

// LogInfo logs messages at the info level
func LogInfo(args ...interface{}) {
	sol.Info(args...)
}

// LogInfof logs messages at the info level
func LogInfof(msg string, args ...interface{}) {
	sol.Infof(msg, args...)
}

// LogTrace logs messages at the trace level
func LogTrace(args ...interface{}) {
	sol.Trace(args...)
}

// LogTracef logs messages at the trace level
func LogTracef(msg string, args ...interface{}) {
	sol.Tracef(msg, args...)
}

// LogWarn logs messages at the warning level
func LogWarn(args ...interface{}) {
	sol.Warn(args...)
}

// LogWarnf logs messages at the warning level
func LogWarnf(msg string, args ...interface{}) {
	sol.Warnf(msg, args...)
}
