// app.go contains all of your application specific settings. Most changes
// needed for a new application should be here, including environment variable
// names, default settings, etc.
package main

import (
	"github.com/mohae/contour"
)

// Name is the name of the application
var Name string = "quine"

// AppCode is the code for the application. This is used to prefix the
// environment variable. It can be left empty.
var AppCode string

// ConfigFil is the configuration file for the application.
var ConfigFilename string = "config.json"

// set-up the application defaults and let contour know about the app.
func init() {
	contour.SetAppCode(AppCode)
	contour.SetConfigFilename(ConfigFilename)

	contour.AddBoolFlag("logging", "l", false,  true, false) 
	contour.AddStringBasic("logconfig", "seelog.xml") 

	InitApp()
}

// set-up custom defaults for your application,
func InitApp() {
	contour.AddBoolBasic("lower", false)
}


