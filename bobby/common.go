// bobby is the sample application. 
package bobby

import (
	"github.com/spf13/viper"
)

// appName is a constant for the application name. It can be useful.
const appName = "quine"

func AppName() string {
	return appName
}

var Config config

type config struct {
	// ConfigFile points to the application's default configuration file.
	// This file can be in TOML, YAML, or JSON format.
	File string

	// ConfigSource are additional paths the application should search for
	// the ConfigFile.
	Source  []string

	// DefaultLogging is the default setting for whether logging is enabled
	// or not.
	DefaultLogging bool

	// DefaultLogFile points to the application's default configuration
	// file for logging. Since seelog is used, a logging configuration file
	// is used too.
	DefaultLogFile string
)

// Setting constants are used for application settings. Its namne is usually
// consistent with its value.
const (
	Config    = "config"
	Logging   = "logging"
	LogConfig = "logconfig"
	Lower     = "lower"
)

func init() {
	Config = &config{ConfigSource: []string{}}
	Config.File = "config.json"
	Config.DefaultLogFile = "seelog.xml"
}

func GetConfig() interface{} {
	return Config
}

// SetConfig loads the applications configuration file and sets the
// application's defaults. The Default* variables have been merged with the
// passed args at this point.
func SetConfig() {
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
	viper.SetDefault(Logging, Config.DefaultLogging)
	viper.SetDefault(LogConfig, Config.DefaultLogFile)
	viper.SetDefault(Config, Config.File)



}


