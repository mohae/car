// hello runs the hello command
package clitpl

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

func Hello(words ...string) (string, error) {
	logger.Tracef("Entering Hello with: %v\n", words)

	h := "Hello"
	if len(words) == 0 {
		logger.Tracef("exiting Hello: h  = %v, err=nil\n", h)
		return h, nil
	}

	for _, word := range words {
		h += " " + word
	}

	if viper.GetBool(Lower) {
		h = strings.ToLower(h)
	}

	// Print out the current settings.
	fmt.Printf("config: %s\n", viper.GetString(Config))
	fmt.Printf("lower: %v\n", viper.GetBool(Lower))
	fmt.Printf("logging: %v\n", viper.GetBool(Logging))
	fmt.Printf("logconfig: %s\n", viper.GetString(LogConfig))

	// This will go to their defined locations
	logger.Trace("This is an example TRACE message\n")
	logger.Debug("This is an example DEBUG message\n")
	logger.Info("This is an example INFO message\n")
	logger.Warn("This is an example WARN message\n")
	logger.Error("This is an example ERROR message\n")
	logger.Critical("This is an example CRITICAL message\n")
	logger.Tracef("exiting Hello\n", h)
	return h, nil
}
