// Contains log related stuff.
package app

import (
	_ "fmt"
	"io/ioutil"

	log "github.com/Sirupsen/logrus"
	"github.com/mohae/contour"
)

func init() {
	//Disable logger by default
	DisableLog()
}

// DisableLog disables all package output
func DisableLog() {
	log.SetOutput(ioutil.Discard)
}

// SetLog sets up logging, if it is enabled to stdout. At this point, the
// only overrides to logging will occur with CLI args. If the CLI args have any
// logging related flags, those will be processed and logging will be updated.
//
func SetLogging() error {
	if !contour.GetBool(EnvLog) {
		// If we're not logging disable the log output
		DisableLog()
		return nil
	}

	log.SetLogLevel(contour.GetString(EnvLogLevel))
	SetAppLogging()
	return nil
}
