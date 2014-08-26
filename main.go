// Copyright © 2014, All rights reserved
// Joel Scoble, https://github.com/mohae/baller
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
// Remember, this application can DELETE stuff. Use at your own risk with a
// clear understanding of what DELETE means. Even though this program will not
// do a delete unless the delete targets have been successfully archived, bugs
// do occur, as do unintended side-effect, and other miscellany that leads to
// problems with software, unfortunately.
//
// CONSIDER YOURSELF WARNED!
//
// Notes on code in Main: some of the code in runMain is copied from the copy-
// right holder, Mitchell Hashimoto (github.com/mitchellh), as I am using his
// cli package.
package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/mitchellh/cli"
	"github.com/mohae/baller/ballr"
	log "github.com/cihub/seelog"
)

func main() {
	// main wraps runMain() and ensures that the log gets flushed prior to exit.
	// Use maxproxs as compression is CPU greedy.
	runtime.GOMAXPROCS(runtime.NumCPU())
	defer logCleanup()
	logging()
	// Exit with return code from runMain()
	rc := runMain()
	os.Exit(rc)
}

func runMain() int {
	log.Info("Baller starting with args ", os.Args[:])
	// runMain parses the flags for logging, sets up CLI stuff for the supported
	// subcommands and runs baller.
	var err error
	err = ballr.SetEnv()
	if err != nil {
		fmt.Println("An error while processing baller Environment variables: ", err.Error())
		return -1
	}

	args := os.Args[1:]

	// Get the command line args. We shortcut "--version" and "-v" to
	// just show the version.
	for _, arg := range args {
		if arg == "-v" || arg == "--version" {
			newArgs := make([]string, len(args)+1)
			newArgs[0] = "version"
			copy(newArgs[1:], args)
			args = newArgs
			break
		}
	}

	cli := &cli.CLI{
		Args:     args,
		Commands: Commands,
		HelpFunc: cli.BasicHelpFunc("baller"),
	}

	exitCode, err := cli.Run()

	if err != nil {
		log.Error("Baller terminating", "error", err.Error())
	}

	log.Info("Baller exiting with exit code ", exitCode)
	return exitCode
}

// logCleanup flushes logs prior to close
func logCleanup() {
	// Do I need to call package's flush log separately? I think log,Flush()
	// should take care of it.
	// TODO find out.
	ballr.FlushLog()
	log.Flush()
}
// logging sets application logging settings.
func logging() error {
	logger, err := log.LoggerFromConfigAsFile("seelog.xml")
	if err != nil {
		return err
	}

	log.ReplaceLogger(logger)

	// setup library loggers
	ballr.UseLogger(logger)

	return nil
}
