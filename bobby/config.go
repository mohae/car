// bobby is the sample application. 
package bobby

import (
	"github.com/spf13/viper"
)

// Not sure if this is needed for handling with ENVs.
// It might be just keeping track of config variables, as constants. is enough.
var Config config

// config is the struct for holding the application config.
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

	// Application specific settings. set the defaults in initAppConfig()
	App	map[string]interface{}
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
//
// SetConfig tries to be as dumb as possible.
//
// First, it checks to see if which of the application's settings already 
// exist. For those that do exist, it notes that fact. Those values can only
// ben overridden if the are flagged as settable and they are passed as 
// supported command-line flags. The already set environment variables are
// flagged as not settable.
// 
// Second, the hard-coded defaults are saved to their respective ENV variables.
//
// Third, the config file settings are read and saved to their respective ENV
// variables, overwriting the application defaults, when applicable.
//
// Finally, the CLI flags are parsed and the relevant, supported flags are 
// saved to their respective ENV variable.
//
// So Precedence, from high to low:
//
//   * CLI args and flags
//   * Environment variable settings.
//   * Application and logging config files
//   * Application defaults.
// 
func SetConfig() {
	// InitConfig save's the application defaults to the env variable.
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


