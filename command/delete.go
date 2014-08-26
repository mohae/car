package command

import (
	"flag"
	"fmt"
	"strings"

	"github.com/mitchellh/cli"
)

// DeleteCommand is a Command implementation that creates an compressed archive
// consisting of the contents of the passed paths and, when it has been 
// successfully created. it deletes the passed paths.
type DeleteCommand struct {
	UI cli.Ui
}

// Help prints the help text for the delete sub-command.
func (c *DeleteCommand) Help() string {
	helpText := `
Usage: baller delete [flags] <archive_name> <paths...>

    $ baller delete dirBackup dir1

The above command creates an archive named dirBackup, using ballers
default comppression algorithm, of dir1. After the archive has been
successfully created, dir1 will be deleted. 

baller will automatically append the correct extension for the
algorithm used if the archive name doesn't already end with the
proper ending.`

	return strings.TrimSpace(helpText)
}

// Run runs the run sub-command; the args are a variadic list of build list names.
func (c *DeleteCommand) Run(args []string) int {

	cmdFlags := flag.NewFlagSet("delete", flag.ContinueOnError)
	cmdFlags.Usage = func() { c.UI.Output(c.Help()) }
//	cmdFlags.StringVar(&logLevel, "log-level", "INFO", "log level")

	fmt.Printf("%+v\n", args)
	fmt.Printf("Not implemented")

	return 0

}

// Synopsis provides a precis of the run sub-command.
func (c *DeleteCommand) Synopsis() string {
	return "Create an archive using the passed paths; delete the passed paths."
}
