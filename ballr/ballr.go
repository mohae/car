// ballr is the main package for baller
package ballr

import (
	"os"

	"github.com/BurntSushi/toml"
	seelog "github.com/cihub/seelog"
)

var logger seelog.LoggerInterface

const (
	EnvBallerTOML	= "EnvBallerTOML"
)

func init() {
	// Disable logger by default
	DisableLog()
}

// AppConfig contains the application configuration settings.
var AppConfig appConfig

type appConfig struct {
	append_datetime bool
	delete bool
	log bool
	datetime_format string
	default_compression string
	default_datetime_prefix string
	default_destination string
}
// DisableLog disables all package log output
func DisableLog() {
	logger = seelog.Disabled
}

// UseLogger uses a specified seelog.LoggerInterface to output lobrary log.
// Use this func if you are using Seelog logging system in your app--I am
// so I'm done
func UseLogger(newLogger seelog.LoggerInterface) {
	logger = newLogger
}

// Call this before app shutdown
func FlushLog() {
	logger.Flush()
}

// SetEnv sets the environment variables, if they do not already exist
func SetEnv() error {
	var err error
	var tmp string
	tmp = os.Getenv(EnvBallerTOML)

	if tmp == "" {
		tmp = "baller.toml"
	}

	_, err  = toml.DecodeFile(tmp, &AppConfig)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	return nil
}
