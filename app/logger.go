// Contains log related stuff.
package app 

import (
	"errors"
	"io"

	seelog "github.com/cihub/seelog"
)

var logger seelog.LoggerInterface

func init() {
	//Disable logger by default
	DisableLog()
}

// DisableLog disables all package output
func DisableLog() {
	logger = seelog.Disabled
}

// UseLoggers uses a specified seelog.LoggerInterface to output package to log.
func UseLogger(newLogger seelog.LoggerInterface) {
	logger = newLogger
	setLibraryLogger()
}

// SetLogWriter uses a specified io.Writer to output library log.
// Use this func if you are not using Seelog logging system in your cmd.
func SetLogWriter(writer io.Writer) error {
	if writer == nil {
		return errors.New("Nil writer")
	}

	newLogger, err := seelog.LoggerFromWriterWithMinLevel(writer, seelog.TraceLvl)
	if err != nil {
		return err
	}

	UseLogger(newLogger)
	setLibraryLogger()
	return nil
}

// FlushLog, call before cmd shutdown. This is called by realMain(). If a
// logger other than Seelog is going to be used, use the 
func FlushLog() {
	flushLibraryLog()
	logger.Flush()
}
