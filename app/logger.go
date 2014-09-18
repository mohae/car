// Contains log related stuff.
package app

import (
	"fmt"
	"os"
	"io"
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

	log.Debugf("finalize logging: should be last entry to temp before copying")
	log.Debugf("Just making sure we are logging to temp. LogFile.Name(): %q", LogFile.Name())
	// See if a logfile is set, if it is, move the temp logfile to the
	// filename and reopen for logging.
	filename := contour.GetString(EnvLogFilename)
	if filename != "" {
		// Make sure its been written to persistent
		err := LogFile.Sync()
		if err != nil {
			log.Fatal(err)
		}

/*
// COPY results in 0 bytes copied, even though there are 21 in LogFile
		logFile, err := os.OpenFile(filename, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
		if err != nil {
			log.Fatal(err)
			return err
		}

		b, err := io.Copy(logFile, LogFile)
		if err != nil {
			log.Fatal(err)
			return err
		}
	
		fmt.Printf("%d copied from %q to %q", b, LogFile.Name(), logFile.Name())
		err = LogFile.Close()
		if err != nil {
			log.Fatal(err)
		}

*/

// Try reading the contents then writing...take slightly longer but what's a
// few milliseconds between friends?
		logFile, err := os.OpenFile(filename, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
		if err != nil {
			fmt.Println(err)
			return err
		}


		tmpLog, err := ioutil.ReadFile(LogFile.Name())
		if err != nil {
			fmt.Println(err)
			return err
		}
	
		
		b, err := logFile.Write(tmpLog)
		if err != nil {	
			fmt.Println(err)
			return err
		}
		fmt.Printf("%d bytes written to %s", b, logFile.Name())

		err = os.Remove(LogFile.Name())
		if err != nil {
			fmt.Println(err)
			return err
		}

		LogFile = logFile
		log.SetOutput(LogFile)
		log.Debugf("Logging to %q", LogFile.Name())
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

	fmt.Println("TempLogFile: ", LogFile.Name())
	log.Debugf("TempLogFile: %s\n", LogFile.Name())

	b, err := io.WriteString(LogFile, "this is a test output\n")
	if err != nil {
		fmt.Fprintf(os.Stderr, "test output err %s", err)
		return err
	}
	fmt.Println(b, "written")
	log.Debugf("This is a test output using log.Debugf")
	return nil
}
