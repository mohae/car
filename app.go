// app.go contains all of your application specific settings. Most changes
// needed for a new application should be here, including environment variable
// names, default settings, etc.
package main

import (

)

// Name is the name of the application
var Name string = "quine"
var configFile string = "config.json"

// appDefaults set the default settings for the application.
appDefaults map[string]setting = make map[string]setting

// setting holds the information for a configuration setting.
type setting struct {
	// Name of the setting
	Name string

	// Type is the datatype for the setting
	Type string

	// The default value of the setting
	Default interface{}

	// IsFlag:  whether or not this is a flag.
	IsFlag bool

	// IsCore; whether or not this is a core setting. If it is, it will
	// override anything set in its ENV variable, regardless of other
	// settings.
	IsCore bool
}


// set-up the application defaults
func init() {
	AddSetting("logging", "bool", false,  true, false) 
	AddSetting("loglevel", "string", "info",  true, false) 
	AddSetting("logging", "string", "error",  true, false) 

	InitApp()
}

// set-up custom defaults for your application,
func InitApp() {
	AddSetting("lower", "bool", false, true, false)
}


