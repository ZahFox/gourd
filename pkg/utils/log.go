package utils

import (
	log "github.com/sirupsen/logrus"
)

var sol *log.Logger = log.New()
var sel *log.Logger = log.New()
var ready = false

// SetupLogging is used to configure the primary loggers
func SetupLogging(fn func(*log.Logger, *log.Logger)) {
	if !ready {
		ready = true
		fn(sol, sel)
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

// LogFatal logs messages at the error level
func LogFatal(args ...interface{}) {
	sel.Fatal(args...)
}

// LogFatalf logs messages at the error level
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
