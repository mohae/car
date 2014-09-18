// Contains log related stuff.
package app

import (
	"fmt"
	"os"
	"io/ioutil"

	log "github.com/Sirupsen/logrus"
	"github.com/mohae/contour"
)

var loggingFinalized bool
var LogFile *os.File

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
	if contour.ConfigProcessed() && !loggingFinalized {
		err := finalizeLogging()
		if err != nil {
			return err
		}
	}
	

	formatter := contour.GetString(EnvLogFormat)

	switch formatter {
	case "json":
		log.SetFormatter(&log.JSONFormatter{})
	default:
		log.SetFormatter(&log.TextFormatter{})
	}

	log.SetLogLevel(contour.GetString(EnvLogLevel))
	return nil
}

// finalize logging is called when all configuration processing has been done. 
// This is in flux because what really needs to be handled is multiwriter output
// support, e.g. stdout and logging or some other destination.
func finalizeLogging() error {
	loggingFinalized = true
	// See about output to logfile:
	if !contour.GetBool(EnvLog) {
		DisableLog()
		return nil
	}

	// See if a logfile is set, if it is, move the temp logfile to the
	// filename and reopen for logging.
	filename := contour.GetString(EnvLogFilename)
	if filename != "" {
		LogFile.Close()
		err := os.Rename(LogFile.Name(), filename)
		if err != nil {
			log.Fatal(err)
			return err
		}
	
		LogFile, err = os.OpenFile(filename, os.O_RDWR | os.O_CREATE, 0666)
		if err != nil {
			log.Fatal(err)
			return err
		}

		log.SetOutput(LogFile)
	}
	
	return nil
}

func SetTempLogFile() error  {
	// First set logging to output to a temp file, this may be moved or
	// deleted when the config and flags get processed.
	var err error
	LogFile, err = ioutil.TempFile("", Name + "-log")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to setup temp logfile: %s", err)
		return err
	}	

	log.SetOutput(LogFile)

	return nil
}
