package app

import (
	"log"
	"time"

	"github.com/mohae/contour"
	jww "github.com/spf13/jwalterweatherman"
)

var (
	// Name is the name of the application
	Name string = "car"

	// ConfigFile is the name of the configuration file for the application.
	CfgFilename string = "app.json"
)

// Variables for configuration entries, or just hard code them.
var (
	CfgFile     string = "cfgfile"     // configuration filename; format type is inferred from ext.
	Log         string = "log"         // to log or not to log
	LogFile     string = "logfile"     // output filename for log output, stderr if empty
	Verbose            = "Verbose"     // Verbose output bool
	VerboseFile        = "VerboseFile" // output filename for Verbose output; stdout if empty.

	Format string = "format" // default archive format
	Type   string = "type"   // default compression type; does not apply to zip archives
)

var unsetTime time.Time

// set-up the application defaults and let contour know about the app.
// Setting also saves them to their relative environment variable.
func init() {
	//
	contour.RegisterCfgFile(CfgFile, CfgFilename)
	contour.SetName(Name)
	contour.SetUseEnv(true)
	contour.SetErrOnMissingCfg(false)
	// Logging and output related
	contour.RegisterBoolFlag(Log, "l", false, "false", "enable/disable logging")
	contour.RegisterBoolFlag(Verbose, "v", false, "false", "Bool for verbose output")
	contour.RegisterStringFlag(Format, "f", "tar", "tar", "create an archive using the tar format")
	contour.RegisterStringFlag(Type, "t", "gzip", "gzip", "create an archive using the zip format")
	// Create Operation Local file selection
	contour.RegisterBoolFlag("delete-files", "D", false, "false", "remove files after adding them to the archive")
	contour.RegisterStringFlag("exclude", "", "", "", "exclude files, given as a PATTERN")
	contour.RegisterStringFlag("exclude-ext", "e", "", "", "exclude files with EXTENSIONS")
	contour.RegisterStringFlag("exclude-anchored", "", "", "", "exclude patterns match file name start")
	contour.RegisterStringFlag("include", "", "", "", "include files, given as a PATTERN")
	contour.RegisterStringFlag("include-ext", "i", "", "", "include files with EXTENSIONS")
	contour.RegisterStringFlag("include-anchored", "", "", "", "include patterns match file name start")
	contour.RegisterStringFlag("outputdir", "O", "", "", "output directory for extract")
	//// Extract Operation Modifiers
	//	contour.RegisterBoolFlag("keep-old-files", "k", false, "don't replace existing files when extracting")
	//	contour.RegisterBoolFlag("keep-newer-files", "", false, "don't replace existing files that are newer than their archive copies")
	//	contour.RegisterBoolFlag("overwrite", "", true, "overwrite existing files when extracting")
	////	contour.RegisterBoolFlag("atime-preserve", "", true)
	//	contour.RegisterBoolFlag("modification-time", "m", false, "don't extract file modified time")
	//	contour.RegisterBoolFlag("same-owner", "", true, "try extracting files with the same ownership")
	//	contour.RegisterBoolFlag("no-same-owner", "", false, "extract files as yourself")
	//	contour.RegisterBoolFlag("numeric-owner", "", true, "always use numbers for user/group names")
	//	contour.RegisterBoolFlag("same-permissions", "p", true, "extract permissions information")
	//	contour.RegisterBoolFlag("no-same-permissions", "", false, "do not extract permissions information")

	//	contour.RegisterBoolFlag("wildcards", "", false, "false", "patterns use wildcards")
	//	contour.RegisterBoolFlag("no-wildcards", "", true, "true", "patters do not use wildcards")
	//	contour.RegisterTimeFlag("newer", "N", unsetTime, "only store files newer than DATE or File")
	//contour.RegisterTimeFlag("newer-mtime", "M", unsetTime, "not set", "only store files modified since DATE")
	//	contour.RegisterStringFlag("newer-file", "", "only store files newere than the DATE for FILENAMEE")

	// Register option aliases
	//	contour.RegisterFlagAlias("newer-date", "after-date")
	//	contour.RegisterFlagAlias("same-permissions", "preserve-permissions")
}

// SetCfg initializes the application's configuration.
func SetCfg() error {
	err := contour.SetCfg()
	if err != nil {
		jww.ERROR.Print(err)
		return err
	}
	SetAppLogging()
	return nil
}

// SetAppLog sets the logger for package loggers and allow for custom-
// ization of the applications log.
func SetAppLogging() {
	if !contour.GetBool(Log) {
		jww.DiscardLogging()
	}
	// set the flags
	jww.TRACE.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	jww.DEBUG.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	jww.INFO.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	jww.WARN.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	jww.ERROR.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	jww.CRITICAL.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	jww.FATAL.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	// get the logfilename, if it's not set, use stderr
	if contour.GetString(LogFile) != "" {
		jww.SetLogFile(contour.GetString(LogFile))
	}
}
