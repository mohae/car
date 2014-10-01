package app

import (
	"time"

	log "github.com/cihub/seelog"
	car "github.com/mohae/carchivum"
	contour "github.com/mohae/contourp"
)

var (
	// Name is the name of the application
	Name string = "car"

	// ConfigFile is the name of the configuration file for the application.
	ConfigFile string = "app.json"

	// LogConfigFile is the name for the log configuration file.
	LogConfigFile string = "seelog.xml"
)

// Variables for configuration entries, or just hard code them.
var (
	CfgConfigFile    string = "configfile"
	CfgLogConfigFile string = "logconfigfile"
	CfgLog           string = "log"
	CfgFormat        string = "format"
	CfgType          string = "type"
)

// Application config.
var Config = contour.AppConfig()
var unsetTime time.Time
// set-up the application defaults and let contour know about the app.
// Setting also saves them to their relative environment variable.
func init() {
	// Calling Register* saves the configuration setting information
	// without writing it to its respective environment variable. This
	// allows any already set environment variables to override non-core
	// vars.
	//
	// Only settings that have been initialized are recognized by contour.

	// The config filename is handled differently, calling this function
	// also sets the ConfigFile format automatically,based on the
	// extension, if it can be determined. If it cannot, the extension is
	// left blank and must be set.
	contour.RegisterConfigFilename(ConfigFile)

	//// Alternative way, manually setting the values
	//contour.RegisterString("configfilename", ConfigFilename)
	//contour.RegisterString("configfileext", "json")

	// Core settings are only settable by the application, and once set are
	// immutable
	contour.RegisterStringCore("name", Name)

	// Immutable settings are only settable once. If its value isn't set
	// during registration, it can be set at a later time. Once it is set,
	// immutable values cannot be changed. Because of this, and the fact
	// that initialization causes bools to be set, bools cannot be made
	// immutable.

	// Set*Flag allows you to add settings that are also exposed as
	// command-line flags. Default implicit values to settings:
	//	IsFlag = true
	//	IsIdempotent = false
	//	IsCore = false
	// The shortcode, 2nd parameter, can be left as an empty string, ""
	// if this flag doesn't support a shortcode.

	// Logging and output related
	contour.RegisterBoolFlag(CfgLog, "l", false, "enable/disable logging")
	contour.RegisterStringFlag(CfgLogConfigFile, "g", LogConfigFile, "name of the log configuration file")

	// AddSettingAlias sets an alias for the setting.
	// contour doesn't support alias yet
	//	contour.AddSettingAlias(CfgLog, "logenabled")
	initApp()

	// Now that the configuration in
}

// InitApp is the best place to add custom defaults for your application,
func initApp() {
	contour.RegisterStringFlag(CfgFormat, "f", "tar", "create an archive using the tar format")
	contour.RegisterStringFlag(CfgType, "t", "gzip", "create an archive using the zip format")

	// Create operation modifiers
	contour.RegisterIntFlag("owner", "", 0, "force UID as owner for added files")
	contour.RegisterIntFlag("group", "", 0, "force GID as group for added files")
	contour.RegisterInt64Flag("mode", "", 0, "force MASK as mode for added files")

	
//	contour.RegisterBoolFlag("usefullpath", "u", false, ")
	
	// Extract Operation Modifiers
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

	// Create Operation Local file selection
	contour.RegisterBoolFlag("delete-files", "D", false, "remove files after adding them to the archive")
	contour.RegisterStringFlag("exclude", "", "", "exclude files, given as a PATTERN")
	contour.RegisterStringFlag("exclude-ext", "e", "", "exclude files with EXTENSIONS")
	contour.RegisterStringFlag("exclude-anchored", "", "", "exclude patterns match file name start")
	contour.RegisterStringFlag("include", "", "", "include files, given as a PATTERN")
	contour.RegisterStringFlag("include-ext", "i", "", "include files with EXTENSIONS")
	contour.RegisterStringFlag("include-anchored", "", "", "include patterns match file name start")
//	contour.RegisterBoolFlag("wildcards", "", false, "patterns use wildcards")
//	contour.RegisterBoolFlag("no-wildcards", "", true, "patters do not use wildcards")
//	contour.RegisterTimeFlag("newer", "N", unsetTime, "only store files newer than DATE or File")
	contour.RegisterTimeFlag("newer-mtime", "M", unsetTime, "only store files modified since DATE")
//	contour.RegisterStringFlag("newer-file", "", "only store files newere than the DATE for FILENAMEE")




	// Register option aliases
	//	contour.RegisterFlagAlias("newer-date", "after-date")
	//	contour.RegisterFlagAlias("same-permissions", "preserve-permissions")
}

// InitConfig initialized the application's configuration. When the config is
// has been initialized, the preset-enivronment variables, application
// defaults, and your application's configuration file, should it have one,
// will all be merged according to the setting properties.
//
// After this, only overrides can occur via command flags.
func SetConfig() error {
	// Set config:
	//    Checks environment variables for settings, follows update rules.
	//    Retrieves config file and applies those settings, if and where
	//      applicable.
	//    Writes the resulting configuration settings to their env vars.
	//  After set config, only command flags can override the settings.
	//  If this is an interactive application, preference changes would
	//    also override certain settings. It may necessitate an additional
	//    flag or two.
	return contour.SetConfig()
}

// SetAppLog sets the logger for package loggers and allow for custom-
// ization of the applications log. This is where app specific code for
// setting up the application's log should be.
//
// SetAppLog assumes that log is enabled if it has been called as its
// caller should be SetLog(). If you are going to call this from elsewhere,
// first make sure that log is enabled.
//
// This uses seelog.
func SetAppLogging() {
	contour.UseLogger(logger)
	car.UseLogger(logger)
	return
}

func DisableAppLogging() {
	contour.DisableLog()
	car.DisableLog()
	return
}

func AppFlushLog() {
	contour.FlushLog()
	car.FlushLog()
	log.Flush()
}
