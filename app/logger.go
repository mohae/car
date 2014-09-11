// Contains log related stuff.
package app

import (
	_"fmt"

	log "github.com/cihub/seelog"
	"github.com/mohae/contour"
)

// constants for loglevels
const (
	Trace    = "trace"
	Debug    = "debug"
	Info     = "info"
	Warn     = "warn"
	Error    = "error"
	Critical = "critical"
	Off      = "off"
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

// SetLog sets up logging, if it is enabled to stdout. At this point, the
// only overrides to logging will occur with CLI args. If the CLI args have any
// logging related flags, those will be processed and logging will be updated.
//
func SetLogging() error {
	if !contour.GetBool(EnvLog) {
		return nil
	}

	logConfig := contour.GetString(EnvLogConfigFile)

	logger, err := log.LoggerFromConfigAsFile(logConfig)
	if err != nil {
		return err
	}
	
	log.ReplaceLogger(logger)
	SetAppLogging()
	return nil
}

func FlushLog() {
	// Flush the library logs.
	AppFlushLog()

	// Then flush the main logger
	logger.Flush()
}

