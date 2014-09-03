// app.go contains all of your application specific settings. Most changes
// needed for a new application should be here, including environment variable
// names, default settings, etc.
package main

import (
	"github.com/mohae/quine/bobby"
	"github.com/mohae/contour"
)

// Name is the name of the application
var Name string = "quine"

// The git commit that was compiled. This will be filled in by the compiler
var GitCommit string

// The main version number that is being run at the moment.
const Version = "0.0.1"

// A pre-release marker for the version. If this is "" (empty string)
// then it means that it is a final release. Otherwise, this is a pre-release
// such as "dev" (in development), "beta", "rc1", etc.
const VersionPrerelease = "dev"

// AppCode is the code for the application. This is used to prefix the
// environment variable. It can be left empty.
var AppCode string

// ConfigFilename is the configuration file for the application.
var ConfigFilename string = "config.json"

// LogConfigFilename is the name for the logging configuration file.
var LogConfigFilename string = "seelog.xml"

// Logging: whether or not application logging is enabled by default.
// Initialize to true if it should automatically be enabled.
var Logging bool

// Environment variables
var (
	EnvConfigFilename string = "configfilename"
	EnvLogConfigFilename string = "logconfigfilename"
	EnvLogging string = "logging"
)

// Config a pointer to the AppConfig. The AppConfig can either be updated by
// calling the contour function or the Config's method, both of which will be
// the same other than being a function or method. 
//
// If you want a different Config object to use for your configuration, call 
// contour.NewConfig() instead. This will return a new Config object. You will
// need to use its methods to work with it, calling contour's function won't 
// apply to it.
var Config *contour.Config = contour.GetAppConfig()

// set-up the application defaults and let contour know about the app.
// Setting also saves them to their relative environment variable.
func init() {
	// Idempotent settings are ones that do one change once they are set.
	// Any subsequent attempts to set an idempotent's setting will not 
	// result in that value being updated.
	// For convenience, each supported datatype can be called either of two
	// ways to make them idempotent. Below is an example for string.
	contour.SetIdempotentString("appname", Name) 
	contour.SetIdemString(EnvConfigFilename, ConfigFilename) 
	contour.SetIdemString(EnvLogConfigFilename, LogConfigFilename)

	// Set*Flag allows you to add settings that are also exposed as
	// command-line flags. Default implicit values to settings:
	//	IsFlag = true
	//	IsIdempotent = false
	// The shortcode, 2nd parameter, can be left as an empty string, ""
	// if this flag doesn't support a shortcode.
	contour.SetBoolFlag(EnvLogging, "l", Logging) 

	// AddSettingAlias sets an alias for the setting.
	contour.AddSettingAlias(EnvLogging, "logenabled")

	InitApp()
}

// InitApp is the best place to add custom defaults for your application,
func InitApp() {
	contour.SetBoolFlag("lower", "", false)
}

// InitConfig initialized the application's configuration. When the config is
// has been initialized, the preset-enivronment variables, application 
// defaults, and your application's configuration file, should it have one,
// will all be merged according to the setting properties.
//
// After this, only overrides can occur via command flags.
func InitConfig() error {
	// Load the already existing environment variables. Only updateable
	// settings are set from these values.
	contour.SetFromEnv()

	// Load the config file
	err := contour.SetFromConfigFile()
	if err != nil {
		return err
	}

	return nil
}

// SetAppLogging sets the logger for package loggers. Any packages that you
// are using that supports logging, configure them here. 
// This uses seelog.
func SetAppLogging() {
	contour.UseLogger(logger)
	bobby.UseLogger(logger)
}
