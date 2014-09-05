package command

import (
	_"fmt"
	"flag"
	"strings"

	"github.com/mohae/cli"
	"github.com/mohae/contour"
	"github.com/mohae/quine/bobby"
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
Usage: quine hello [flags] <wordlist string...>

hello will take a 1 or more words, concatonate them and print the resulting
string to stdout.

    $ quine hello how are you
    Hello how are you

    $ quine -lower hello world
    Hello world

quine flags supports flags that most any application would have, along
with one example flag. 

    --lower=(true, false)    true lowercases the output.

    --logging=(true, false)  enable/disable log output
    -l                       alias to --logging
`
	return strings.TrimSpace(helpText)
}

// Run runs the hello command; the args are a variadic list of words
// to append to hello.
func (c *HelloCommand) Run(args []string) int {
	// set up the command flags
	cmdFlags := flag.NewFlagSet("run", flag.ContinueOnError)
	cmdFlags.Usage = func() {
		c.UI.Output(c.Help())
	}

	// Filter the flags from the args and update the config with them.
	// The args remaining after being filtered are returned.
	filteredArgs, err := contour.FilterArgs(cmdFlags, args)
	if err != nil {
		c.UI.Error(err.Error())
		return 1
	}

	// Run the command in the package.
	bobby.Hello(filteredArgs...)

	c.UI.Output("quine Hello is complete")
	return 0
}

// Synopsis provides a precis of the hello command.
func (c *HelloCommand) Synopsis() string {
	ret := `Concatonates the list of words it recieved
to 'Hello', applies any formatting required,
and returns the result.
`

	return ret
}
