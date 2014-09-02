packager main

import (
	"flag"
	
	log "github.com/cihub/seelog"
	"github.com/mohae/contour"
)

// InitConfig Sets the application defaults. This defines the application
// behavior if the config file is missing; if running without a config file
// is going to be supported.
func InitConfig() {
	//dflts := make(map[string]interface{})

	if dflts == nil {
		return
	}
	
	
}

// SetConfigFromFile loads the config file. For each key, it validates that the
// key is an application recognized key and it is settable via the config file.
// If both are true, The settings in this file override the application 
// defaults, when allowed, but not their ENV variable values. The settings are 
// saved to their respective ENV variables.
func SetConfigFromFile() error {
	
	return nil
}

// SetEnv checks to see if the received ENV variable exists. If it does, it
// saves the new value if override of it is allowed.
func SetEnv() error {

	return nil
}

