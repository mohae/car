// hello runs the hello command
package bobby

import (
	"fmt"
	"os"
	"strings"
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

	fmt.Println(h)
	
	return h, nil
}
