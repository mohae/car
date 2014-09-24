// Contains log related stuff.
package app

import (
	log "github.com/cihub/seelog"
	"github.com/mohae/contour"
)

var loggingFinalized bool
var logger log.LoggerInterface
//var LogFile *os.File

func init() {
	//Disable logger by default
	DisableLog()
}

// DisableLog disables all package output
func DisableLog() {
	logger = log.Disabled
	DisableAppLogging()
}

// SetLog sets up logging, if it is enabled to stdout. At this point, the
// only overrides to logging will occur with CLI args. If the CLI args have any
// logging related flags, those will be processed and logging will be updated.
//
func SetLogging() error {
	var err error

	if !contour.GetBool(CfgLog) {
		DisableLog()
		return nil
	}

	logger, err = log.LoggerFromConfigAsFile(contour.GetString(CfgLogConfigFile))
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
