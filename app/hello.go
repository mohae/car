package app

import (
	"fmt"
	"os"
	"strings"
)

func Hello(words ...string) (string, error) {
	h := "Hello"
	if len(words) == 0 {
		
		logger.Tracef("exiting Hello: h  = %v, err=nil\n", h)
		return h, nil
	}

	for _, word := range words {
		h += " " + word
	}

	if os.Getenv(EnvLower) == "true" {
		h = strings.ToLower(h)
	}

	// Print out the current settings.
	fmt.Printf("configfile: %s\n", os.Getenv(EnvConfigFile))
	fmt.Printf("logging: %v\n", os.Getenv(EnvLogging))
	fmt.Printf("logconfigfile: %s\n", os.Getenv(EnvLogConfigFile))
	fmt.Printf("lower: %v\n", os.Getenv(EnvLower))

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

