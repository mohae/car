package app

import (
	"github.com/mohae/contour"
)

var (
	// Name is the name of the application
	Name string = "car"

	// ConfigFile is the name of the configuration file for the application.
	ConfigFilename string = "app.json"

	// Logfile is the name of the default log file for the application
	LogFilename string = "app.log"
)

// Environment variables
var (
	EnvConfigFilename      string = "configfile"
	EnvLogFormat	string = "logformat"
	EnvLogFilename	string = "logfile"
	EnvLog             string = "log"
	EnvLogLevel        string = "loglevel"
	EnvStdout          string = "stdout"
	EnvStdoutLevel     string = "stdoutlevel"
	EnvArchiveFormat   string = "archiveformat"
	EnvCompressionType string = "compressiontype"
	EnvVerbose         string = "verbose"
)

// Application config.
var Config = contour.AppConfig()

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
	contour.RegisterConfigFilename(EnvConfigFilename, ConfigFilename)

	//// Alternative way, manually setting the values
	//contour.RegisterString("configfilename", ConfigFilename)
	//contour.RegisterString("configfileext", "json")

	// Core settings are only settable by the application, and once set are
	// immutable
	contour.RegisterCoreString("appname", Name)

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
	contour.RegisterFlagBool(EnvLog, false, "l")
	contour.RegisterFlagString(EnvLogFilename, LogFilename, "f")
	contour.RegisterFlagString(EnvLogLevel, "warn", "")
	contour.RegisterFlagBool(EnvStdout, false, "s")
	contour.RegisterFlagString(EnvStdoutLevel, "info", "")
	contour.RegisterFlagBool(EnvVerbose, false, "v")

	// AddSettingAlias sets an alias for the setting.
	// contour doesn't support alias yet
	//	contour.AddSettingAlias(EnvLog, "logenabled")

	initApp()

	// Now that the configuration in
}

// InitApp is the best place to add custom defaults for your application,
func initApp() {
	contour.RegisterFlagString("archiveformat", "tar", "f")
	contour.RegisterFlagString("compressiontype", "gzip", "t")
}

// InitConfig initialized the application's configuration. When the config is
// has been initialized, the preset-enivronment variables, application
// defaults, and your application's configuration file, should it have one,
// will all be merged according to the setting properties.
//
// After this, only overrides can occur via command flags.
func InitConfig() error {
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
