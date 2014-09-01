package command

import (
	"flag"
	"fmt"
	"strings"

	"github.com/mitchellh/cli"
	"github.com/mohae/quine/bobby"
	"github.com/spf13/viper"
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
	// Create a new commandflag set.
	cmdFlags = flag.NewFlagSet("hello", flag.ContinueOnError)
	cmdFlags.Usage = func() {
		c.UI.Output(c.Help())
	}

	// Set the filters common to more than one command.
	setCommonFilters()

	// Set the command specific filters here.
	cmdFlags.BoolVar(&lowerFilter, "lower", false, "lower filter")

	// Parse the args for flags.
	err := cmdFlags.Parse(args)
	if err != nil {
		c.UI.Error(fmt.Sprintf("parse of command-line arguments failed: %s", err))
		return 1
	}

	// Set the application configuration
	appConfig()

	// Set the command level configuration.
	viper.Set(clitpl.Lower, lowerFilter)

	// TODO setup logging
	err = logConfig()
	if err != nil {
		c.UI.Error(fmt.Sprintf("setup and configuration of application logging failed: %s", err))
		return 1
	}

	//use the filteredArgs
	// The remaining flags are the words to append. We capture them in a
	// []string in case there is some other processing.or processing is
	// conditional on the existence of args, Hello doesn't care about that.
	helloArgs := cmdFlags.Args()

	// Run the Hello command, passing the Args.
	message, err := clitpl.Hello(helloArgs...)
	if err != nil {
		c.UI.Error(err.Error())
		return 1
	}

	// Print out the returned message.
	c.UI.Output(message)

	c.UI.Output("clitemplate Hello is complete")
	return 0
}

// Synopsis provides a precis of the hello command.
func (c *HelloCommand) Synopsis() string {
	return `Concatonates the list of words it recieved
               to 'Hello', applies any formatting required,
               and returns the result.`
}
