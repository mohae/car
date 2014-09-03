// Contains log related stuff.
package main

import (
	"os"

	log "github.com/cihub/seelog"
	utils "github.com/mohae/utilitybelt"
)

var logger log.LoggerInterface

func init() {
	//Disable logger by default
	DisableLog()
}

// DisableLog disables all package output
func DisableLog() {
	logger = log.Disabled
}

// UseLoggers uses a specified seelog.LoggerInterface to output package to log.
func UseLogger(newLogger log.LoggerInterface) {
	logger = newLogger
}

// FlushLog, call before app shutdown.
// With the way this application is structured, logging needs to be set and
// flushed at the command/name.go level. If you aren't going to support
func FlushLog() {
	logger.Flush()
}

// SetLogging sets up logging, if it is enabled to stdout. At this point, the
// only overrides to logging will occur with CLI args. If the CLI args have any
// logging related flags, those will be processed and logging will be updated.
//
// Notes: this is temporary until the following is done and enabled:
// 	If debug enabled, log to stdout until cli args are processed
//	always log to a tmp file until cli args are processed
//		either the tmp entries get written to the log
//		or tmp gets discarded (if logging disabled)
//	debug enabling will only enable logging output to stdout using the desired level.
func SetLogging() error {
	// If it can't be interpreted as a bool, assume its false.
	b := os.Getenv(AppCode + EnvLogging)
	if !utils.StringIsBool(b) {
		return nil
	}
	
	configFilename := os.Getenv(AppCode + EnvLogConfigFilename)
	if configFilename == "" {
		configFilename = LogConfigFilename		
	}

	logger, err := log.LoggerFromConfigAsFile(configFilename)
	if err != nil {
		return err
	}
	
	log.ReplaceLogger(logger)
	return nil
}

// also be written to temp, until it is determined if logging is enabled or not.

