package command

import (
	"strings"

	"github.com/mohae/car/app"
	"github.com/mohae/cli"
	"github.com/mohae/contour"
)

// ExtractCommand is a Command implementation that says hello world
type ExtractCommand struct {
	UI cli.Ui
}

// Help prints the help text for the run sub-command.
func (e *ExtractCommand) Help() string {
	helpText := `
Usage: car extract [flags] <source> <destination>

This will extract the source to the destination.

The compression type can either be specified or detected by car.

Extract supports the following flags(Type):

    --logging(bool)     enable/disable log output
    --type              compression type of the archive
    --format            the archive format to use, tar
                        is the default, --type is ignored
			when the format is zip, since it
			comes with its own compression type.
    --verbose           verbose output.
`
	return strings.TrimSpace(helpText)
}

// Run runs the extract command; the args are a variadic list of words
// to append to extract.
func (c *ExtractCommand) Run(args []string) int {
	// set up the command flags
	contour.SetUsage(func() {
		c.UI.Output(c.Help())
	})
	// Filter the flags from the args and update the config with them.
	// The args remaining after being filtered are returned.
	filteredArgs, err := contour.FilterArgs(args)
	if err != nil {
		c.UI.Error(err.Error())
		return 1
	}
	app.SetAppLogging()
	if len(filteredArgs) == 0 {
		c.UI.Error("Error: to extract an archive, a source must be specified")
		return 1
	}
	// it is assumed that the first arg is the source, anything else is ignored
	message, err := app.Extract(filteredArgs[0])
	if err != nil {
		c.UI.Error(err.Error())
		return 1
	}
	c.UI.Output(message)
	return 0
}

// Synopsis provides a precis of the hello command.
func (c *ExtractCommand) Synopsis() string {
	ret := `creates a compressed archive from the specified source(s) and writes it out to the destination.`
	return ret
}
