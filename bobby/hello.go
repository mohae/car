// hello runs the hello command
package bobby

import (
	"fmt"
	"os"
	"strings"

	"github.com/mohae/contour"
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

	if os.Getenv("lower") == "true" {
		h = strings.ToLower(h)
	}

	// Print out the current settings.
	fmt.Printf("config: %s\n", contour.GetString("configfilename"))
	fmt.Printf("lower: %v\n", contour.GetBool("lower"))
	fmt.Printf("logging: %v\n", contour.GetBool("logging"))
	fmt.Printf("logconfig: %s\n", contour.GetString("logconfigfilename"))

        fmt.Printf("config: %s\n", os.Getenv("configfilename"))
        fmt.Printf("lower: %v\n", os.Getenv("lower"))
        fmt.Printf("logging: %v\n", os.Getenv("logging"))
        fmt.Printf("logconfig: %s\n", os.Getenv("logconfigfilename"))


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
