package app

import (
	"fmt"
	"strings"

	log "github.com/cihub/seelog"
	"github.com/mohae/contour"
)

func Hello(words ...string) (string, error) {
	h := "Hello"
	if len(words) == 0 {
		
		log.Tracef("exiting Hello: h  = %v, err=nil\n", h)
		return h, nil
	}

	for _, word := range words {
		h += " " + word
	}

	if contour.GetBool(EnvLower) {
		h = strings.ToLower(h)
	}

	fmt.Println(h)

	// This will go to their defined locations
	log.Trace("This is an example TRACE message\n")
	log.Debug("This is an example DEBUG message\n")
	log.Info("This is an example INFO message\n")
	log.Warn("This is an example WARN message\n")
	log.Error("This is an example ERROR message\n")
	log.Critical("This is an example CRITICAL message\n")

	log.Tracef("exiting Hello\n")
	return h, nil
}

