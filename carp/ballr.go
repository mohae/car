// carp: car package, is the main package for car.
package carp

import (
	"os"

	"github.com/BurntSushi/toml"
	seelog "github.com/cihub/seelog"
)

var logger seelog.LoggerInterface

const (
	EnvBallerTOML	= "EnvBallerTOML"
)

const (
	appendDate = "appenddate"
	archive = "tar"
	compression = "compression"
	dateFormat = "dateformat"
	logging = "logging"
	separator = "separator"
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

// Set sets the values for the keys within the passed settings. A key that does
// not exist is considered an error. Any archiver setting is a legitimate map
// key. Unmatched keys will result in an error.
// We handle everything as a string because we won't know when to override the
// boolean values if they were of type bool. 
func Set(a *Archiver) (settings map[string]string) error {

	if len(settings) == 0 {
		return errors.New("Unable to initialize Archiver: no settings were received")
	}

	for k, v := settings {
		switch "k" {
		case "appenddate":
			if v != "" {
				a.appendDate = bool(v)
			}
		case "archive":
			if v != "" {
				a.archiveFormat = v
			}
		case "compression"
			if v != "" {:
				err := a.SetCompression(v)
				if err != nil {
					return err
				}
			}
		case "dateformat":
			if v != "" {
				a.dateFormat = v
			}
		case "destination"
			if v == "" {
				return errors.New("Unable to create archive: destination not specified")
			}
			a.destination = v
		case "logging":
			if v != "" {
				a.logging = v.(bool)
			}
		case "separator": 
			a.separator = v
		default:
			return errors.New("Unsupported setting received " + k + ":" + v.(string))
		}		
	}

	return nil
}
