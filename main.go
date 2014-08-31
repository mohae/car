// Copyright © 2014, All rights reserved
// Joel Scoble, https://github.com/mohae/clitpl
//
// This is licensed under The MIT License. Please refer to the included
// LICENSE file for more information. If the LICENSE file has not been
// included, please refer to the url above.
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License
//
// clitpl is a basic implementation of michellh's cli package. clitpl uses
// cli's example, and the implementations used in both mitchellh's and
// hashicorp's implemenations.
//
package main

import (
	"fmt"
	_ "io/ioutil"
	"os"

	log "github.com/cihub/seelog"
	"github.com/mitchellh/cli"
)

// name is the name of the application.
var name string

// version is the version of the application
var version = Version()

type appInfo struct {
	// Name
	Name	string
	
	// Version
	Version string

	// Commands
	Commands map[string]cli.CommandFactory
	// Args
	Args []string
}

var AppInfo = NewAppInfo()

func NewAppInfo() *appInfo {
	inf := &appInfo{Name: name, Version: version, Commands}
	return inf
}

// This is modeled on mitchellh's realmain wrapper
func main() {
	os.Exit(realMain())
}

// realMain, it's the real main, yo
// no logging is done until the flags are processed, since the flags could
// enable/disable output, alter it, or alter its output locations. Everything
// must go to stdout until then.
func realMain() int {
	defer log.Flush()
	defer clitpl.FlushLog()

	// Get the command line args. We shortcut "--version" and "-v" to
	// just show the version.
	appInfo.Args := os.Args[1:]

	// Setup the args, Commands, and Help info.
	cli := &cli.CLI{
		AppInfo: appInfo,
		HelpFunc: cli.BasicHelpFunc("clitemplate"),
	}

	// Run the passed command, recieve back a message and error object.
	exitCode, err := cli.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error executing CLI: %s\n", err.Error())
		return 1
	}

	// Return the message.
	return exitCode
}

// Version returns the version number.
func Version() string {
	var v string
	return v
}
