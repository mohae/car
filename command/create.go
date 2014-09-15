package command

import (
	"flag"
	"fmt"
	"strings"
<<<<<<< HEAD

	"github.com/mohae/car/app"
	"github.com/mohae/cli"
	"github.com/mohae/contour"
)

// CreateCommand is a Command implementation that says hello world
=======
	"time"

	"github.com/mitchellh/cli"
	"github.com/mohae/carchivum"
)

var archiver carchivum.*Archiver

// CreateFilter has all the valid commandline args for the Create command
type CreateFilter struct {
	// appendDatetime is a boolean for whether to automatically append
	// the current datetime, using the configured datetime format, when
	// there is a collision on the archive name.
	// Notes:
	//   * This only applies to named operations, like files. 
	//   * If the filename format includes the datetime, this setting has
	//     no effect.
	//   * Collision on dateformatted names, including auto-generated names
	//     result in an error
	// TODO add auto-append random number on collision
	appendDate string

	// compressionFormat is the compression format to use for this
	// operation
	compressionFormat string

	// dateFormat is the Go datetime format string for datetime values.
	dateFormat  string

	// logging enables/disables logging
	logging string

	// separator is the separator to use when datetime formats have
	// spaces in them; this value replaces the separator.
	separator string
}


// CreateCommand is a Command implementation that creates an compressed archive
// consisting of the contents of the passed paths.
>>>>>>> 2e926bf1982fb46d62edad497bf67f5e2bff12e7
type CreateCommand struct {
	UI cli.Ui
}

<<<<<<< HEAD
// Help prints the help text for the run sub-command.
func (c *CreateCommand) Help() string {
	helpText := `
Usage: car create [flags] <destination> <source...>

This will create a compressed archive from the list of
sources and write it to the destination.

create supports the following flags(Type):

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

// Run runs the hello command; the args are a variadic list of words
// to append to hello.
func (c *CreateCommand) Run(args []string) int {
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

	err = app.SetLogging()
	if err != nil {
		c.UI.Error(fmt.Sprintf("setup and configuration of application logging failed: %s", err))
		return 1
	}

	// Run the command in the package.
	message, err := app.Create(filteredArgs...)
	if err != nil {
		c.UI.Error(err.Error())
		return 1
	}

	c.UI.Output(message)
	return 0
}

// Synopsis provides a precis of the hello command.
func (c *CreateCommand) Synopsis() string {
	ret := `creates a compressed archive from the specified source(s) and writes it out to the destination."
`

	return ret
=======
// Help prints the help text for the delete sub-command.
func (c *CreateCommand) Help() string {
	helpText := `
Usage: baller create [options] destination <...sources>

    $ baller create dirBackup dir1

The above command creates an archive named dirBackup, using ballers
default comppression algorithm, of dir1. 

baller will automatically append the correct extension for the
algorithm used if the archive name doesn't already end with the
proper ending.

For all options, an absence of the setting will result in the current values
being used.

Options:

-appenddate=true, false    * When true, if there is a collision on the file-
                           name, the current datetime will be appended to the 
                           filename using the current datetime format. 
                           * When false, if there is a collision on the file-
                           name, an error will be returned.

-archive=tar, zip          The format of the archive to be created. Tar is the
                           default. Zip is the only supported alternative. If
                           using the zip format, the compression option will be
                           ignored as zip is a compressed archive format.

-compression=<format name> Use the named compression format instead of the
                           current compression format.

-dateformat=<format>       Set the datetime format for datetime output. Uses 
                           Go's datetime formatting.

-logging=true, false       Controls whether to enable logging or not,

-separator=<string>        The string used to replace spaces in datetime 
                           formats that have them. An empty value removes the
                           spaces.`

	return strings.TrimSpace(helpText)
}

// Run runs the run sub-command; the args are a variadic list of build list names.
func (c *CreateCommand) Run(args []string) int {

	cmdFlags := flag.NewFlagSet("create", flag.ContinueOnError)
	cmdFlags.Usage = func() {
		c.UI.Output(c.Help()) 
	}

	cmdFlags.StringVar(&appendDateFilter, appendDate, "", "appenddate filter")
	cmdFlags.SrtingVar(&archiveFilter, archive, "", "archive filter")
	cmdFlags.StringVar(&compressionFilter, compression, "", "compression filter")
	cmdFlags.StringVar(&dateFormatFilter, dateFormat, "", "dateformat filter")
	cmdFlags.StringVar(&loggingFilter, logging, "", "logging filter")
	cmdFlags.StringVar(&separatorFilter, separator, "", "separator filter")
//	.Printf("%+v\n", args)

	// populate the settings map
	var settings  map[string]string
	if appendDateFilter != "" {
		settings[appendDate] = appendDateFilter
	}
	
	if archiveFilter != "" {
		settings[archive] = archiveFilter
	}

	if compressionFilter != "" {
		settings[compression] = compressionFilter
	}

	if dateFormatFilter != "" {
		settings[dateFormat] = dateFormatFilter
	}

	if loggingFilter != "" {
		settings[logging] = loggingFilter
	}

	if separatorFilter != "" {
		settings[separator] = separatorFilter
	}

	// Parse the passed args for flags,
	if err := cmdFlags.Parse(args); err != nil {
		c.UI.Error(fmt.Sprintf("Parse of command-line arguments failed: %s", err))
		return 1
	}

	// The rest is the filename and sources
	createArgs  := cmdFlags.Args()

	if 
        // Process the request; we don't bother checking the source args since
	// the called function checks it anyways. 
	var message string
	var err error
	message, err = ballr.Create(cmdFlags.Args()...)
	if err != nil {
		c.UI.Error(err.Error())
	}

	c.UI.Output(message)

	return 0

}

// Create is the master process for creating archives.
func Create(args CreateFilter) error {
	switch len(args) {
	case == 0:
		return errors.New("Nothing was received with the `baller create` command, an output and a source, a path, is required at the minimum")
	case == 1:
		return errors.New("No source path was received. At least one source must be included to create an archive.")
	}

	var settings map[string]interface{}

	// The first arg is the destination:
	settings["destination"] = args[0]


}


// Synopsis provides a precis of the run sub-command.
func (c *CreateCommand) Synopsis() string {
	return "Create an archive using the passed paths."

>>>>>>> 2e926bf1982fb46d62edad497bf67f5e2bff12e7
}
