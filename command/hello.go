package command

import (
	_ "flag"
	_ "fmt"
	"strings"

	"github.com/mitchellh/cli"
	_ "github.com/mohae/contour"
	_ "github.com/mohae/quine/bobby"
)

// Hello command specific filter.
var lowerFilter bool

// HelloCommand is a Command implementation that says hello world
type HelloCommand struct {
	UI cli.Ui
}

// Help prints the help text for the run sub-command.
func (c *HelloCommand) Help() string {
	helpText := `
Usage: clitemplate hello [flags] <wordlist string...>

hello will take a 1 or more words, concatonate them and print the resulting
string to stdout.

	$ clitemplate hello how are you
	Hello how are you

	$ clitemplate -lower hello world
	Hello world

clitemplate flags supports flags that most any application would have, along
with one example flag. 

	--lower=(true, false)    true lowercases the output.
        -w                       alias to --lowercase

	--logging=(true, false)  enable/disable log output
        -l                       alias to --logging

	--logconfig=filename     the path to the log configuration file this 
                                 run should use, instead of the application's.
                                 This setting enables logging, regardless of 
                                 any other setting.
        -g                       Alias to --logconfig
`

	return strings.TrimSpace(helpText)
}

// Run runs the hello command; the args are a variadic list of words
// to append to hello.
func (c *HelloCommand) Run(args []string) int {

	c.UI.Output("clitemplate Hello is complete")
	return 0
}

// Synopsis provides a precis of the hello command.
func (c *HelloCommand) Synopsis() string {
	return `Concatonates the list of words it recieved
               to 'Hello', applies any formatting required,
               and returns the result.`
}
