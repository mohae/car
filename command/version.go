package command

import (
	"bytes"
	"fmt"

<<<<<<< HEAD
	"github.com/mohae/cli"
=======
	"github.com/mitchellh/cli"
>>>>>>> 2e926bf1982fb46d62edad497bf67f5e2bff12e7
)

// VersionCommand is a Command implementation that prints the version.
type VersionCommand struct {
<<<<<<< HEAD
	Name		string
=======
>>>>>>> 2e926bf1982fb46d62edad497bf67f5e2bff12e7
	Revision          string
	Version           string
	VersionPrerelease string
	UI                cli.Ui
}

// Help prints the Help text for the version sub-command
func (c *VersionCommand) Help() string {
<<<<<<< HEAD
	return "Prints " + c.Name + "'s version information."
=======
	return "Prints baller's version information."
>>>>>>> 2e926bf1982fb46d62edad497bf67f5e2bff12e7
}

// Run runs the version sub-command.
func (c *VersionCommand) Run(_ []string) int {
	var versionString bytes.Buffer
<<<<<<< HEAD
	fmt.Fprintf(&versionString, "%s v%s", c.Name, c.Version)
=======
	fmt.Fprintf(&versionString, "Rancher v%s", c.Version)
>>>>>>> 2e926bf1982fb46d62edad497bf67f5e2bff12e7
	if c.VersionPrerelease != "" {
		fmt.Fprintf(&versionString, ".%s", c.VersionPrerelease)

		if c.Revision != "" {
			fmt.Fprintf(&versionString, " (%s)", c.Revision)
		}
	}

	c.UI.Output(versionString.String())

	return 0
}

// Synopsis provides a precis of the version sub-command.
func (c *VersionCommand) Synopsis() string {
<<<<<<< HEAD
	return "Prints the " + c.Name + " version"
=======
	return "Prints the baller version"
>>>>>>> 2e926bf1982fb46d62edad497bf67f5e2bff12e7
}
