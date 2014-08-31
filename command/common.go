// common: common variables and functionality for cli command support.
// Anything that is used by more than one command should be put in here for
// codes sake.
package command

import (
	"flag"

	log "github.com/cihub/seelog"
	"github.com/mohae/clitemplate/clitpl"
	"github.com/spf13/viper"
)

var settingFilter, logConfigFilter string
var loggingFilter bool
var cmdFlags *flag.FlagSet

// setCommonFilters is a place to set filters that are common to more than one
// command.
func setCommonFilters() {
	// Set up the argument filters for the flags.
	cmdFlags.BoolVar(&loggingFilter, clitpl.Logging, clitpl.DefaultLogging, clitpl.Logging+" filter")
	cmdFlags.StringVar(&logConfigFilter, clitpl.LogConfig, clitpl.DefaultLogConfig, clitpl.LogConfig+" filter")
}

func mergeCommonFlags() {

}

func appConfig() error {
	// set the passed values
	clitpl.DefaultLogging = loggingFilter
	clitpl.DefaultLogConfig = logConfigFilter

	// Initialize the application's configuration.
	clitpl.InitConfig()

	// Set any additional flags that aren't part of InitConfig()
	// TODO change this to something your application uses, or elide it.
	viper.Set("setting", settingFilter)

	return nil
}

// setupLogging is only done if logging is enabled.
func logConfig() error {
	if !viper.GetBool(clitpl.Logging) {
		return nil
	}

	logger, err := log.LoggerFromConfigAsFile(viper.GetString(clitpl.LogConfig))
	if err != nil {
		return err
	}

	log.ReplaceLogger(logger)
	clitpl.UseLogger(logger)

	return nil
}
