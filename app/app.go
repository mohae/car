package app

import (
	"github.com/mohae/contour"
)

// Name is the name of the application
var Name string = "quine"

// AppCode is the code for the application. This is used to prefix the
// environment variable. It can be left empty.
var AppCode string

// ConfigFile is the name of the configuration file for the application.
var ConfigFile string = "config.json"

// LogConfigFile is the name for the log configuration file.
var LogConfigFile string = "seelog.xml"

// Log: whether or not application log is enabled by default.
// Initialize to true if it should automatically be enabled.
var Log bool

// Environment variables
var (
	EnvConfigFile string = "configfile"
	EnvLogConfigFile string = "logconfigfile"
	EnvLog string = "log"

	EnvLower string = "lower"
)

// Application config.
var Config *contour.Config = contour.AppConfig()

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
	contour.RegisterConfigFilename(EnvConfigFile, ConfigFile)

	//// Alternative way, manually setting the values
	//contour.RegisterString(EnvConfigFilename, ConfigFilename) 
	//contour.RegisterString(EnvConfigFileExt, "json") 

	// Core settings are only settable by the application, and once set are
	// immutable
	contour.RegisterCoreString("appname", Name) 

	// Immutable settings are only settable once. If its value isn't set
	// during registration, it can be set at a later time. Once it is set,
	// immutable values cannot be changed. Because of this, and the fact
	// that initialization causes bools to be set, bools cannot be made
	// immutable.

	// This is set in the config file.
	contour.RegisterImmutableString(EnvLogConfigFile, LogConfigFile)
	
	// Set*Flag allows you to add settings that are also exposed as
	// command-line flags. Default implicit values to settings:
	//	IsFlag = true
	//	IsIdempotent = false
	//	IsCore = false
	// The shortcode, 2nd parameter, can be left as an empty string, ""
	// if this flag doesn't support a shortcode.
	contour.RegisterBoolFlag(EnvLog, Log, "") 

	// AddSettingAlias sets an alias for the setting.
	// contour doesn't support alias yet
	contour.AddSettingAlias(EnvLog, "logenabled")

	initApp()

	// Now that the configuration in
}

// InitApp is the best place to add custom defaults for your application,
func initApp() {
	contour.RegisterBoolFlag(EnvLower, false, "")
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
	return 
}

func AppFlushLog() {
	contour.FlushLog()
}
