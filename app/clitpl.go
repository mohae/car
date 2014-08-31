// clitemplate is the main package for clitemplate.
package clitpl

import (
	_ "errors"
	_ "io"

	"github.com/spf13/viper"
)

// AppName is a constant for the application name. It can be useful.
var AppName = "clitemplate"

var (
	// ConfigFile points to the application's default configuration file.
	// This file can be in TOML, YAML, or JSON format.
	ConfigFile = "config.toml"

	// DefaultLogging is the default setting for whether logging is enabled
	// or not.
	DefaultLogging = false

	// DefaultLogConfig points to the application's default configuration
	// file for logging. Since seelog is used, a logging configuration file
	// is used too.
	DefaultLogConfig = "seelog.xml"
)

// Setting constants are used for application settings. Its namne is usually
// consistent with its value.
const (
	Config    = "config"
	Logging   = "logging"
	LogConfig = "logconfig"
	Lower     = "lower"
)

// InitConfig loads the applications configuration file and sets the
// application's defaults. The Default* variables have been merged with the
// passed args at this point.
func InitConfig() {
	// Configure viper
	viper.SetConfigFile(ConfigFile)

	// Viper also supports adding search paths, call as many times as needed.
	//viper.AddConfigPath(Source)

	// And load the application config file.
	err := viper.ReadInConfig()
	if err != nil {
		logger.Warn("Config not found. The application defaults will be used. This may result in unexpected behavior")
	}

	// viper is case insensitive, though it doesn't matter here since we
	// are using constants. Just wanted to say...
	viper.SetDefault(Logging, DefaultLogging)
	viper.SetDefault(LogConfig, DefaultLogConfig)
	viper.SetDefault(Config, ConfigFile)

}
