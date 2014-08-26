package command

import (
	"flag"
	"fmt"
	"strings"

	"github.com/mitchellh/cli"
)

// CreateCommand is a Command implementation that creates an compressed archive
// consisting of the contents of the passed paths.
type CreateCommand struct {
	UI cli.Ui
}

// Help prints the help text for the delete sub-command.
func (c *CreateCommand) Help() string {
	helpText := `
Usage: baller create [flags] <archive_name> <paths...>

    $ baller create dirBackup dir1

The above command creates an archive named dirBackup, using ballers
default comppression algorithm, of dir1. 

baller will automatically append the correct extension for the
algorithm used if the archive name doesn't already end with the
proper ending.`

	return strings.TrimSpace(helpText)
}

// Run runs the run sub-command; the args are a variadic list of build list names.
func (c *CreateCommand) Run(args []string) int {

	cmdFlags := flag.NewFlagSet("delete", flag.ContinueOnError)
	cmdFlags.Usage = func() { c.UI.Output(c.Help()) }
//	cmdFlags.StringVar(&logLevel, "log-level", "INFO", "log level")

	fmt.Printf("%+v\n", args)
	fmt.Printf("Not implemented")

	return 0

}

// Synopsis provides a precis of the run sub-command.
func (c *CreateCommand) Synopsis() string {
	return "Create an archive using the passed paths."

}
