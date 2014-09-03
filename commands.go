// Initializes the Commands struct for the application.
// New commands need to be added to the CommandFactory map.
package main

import (
	"os"

	"github.com/mohae/cli"
	"github.com/mohae/quine/command"
)

// Commands
var Commands map[string]cli.CommandFactory

// Set-up the commands for the application. Help and version doesn't need to bo
// set-up because they are always available.
func init() {
	ui := &cli.BasicUi{Writer: os.Stdout}
	Commands = map[string]cli.CommandFactory{
		"hello": func() (cli.Command, error) {
			return &command.HelloCommand{
				UI: ui,
			}, nil
		},
		"version": func() (cli.Command, error) {
			return &command.VersionCommand{
				Name: Name,
				Revision: GitCommit,
				Version: Version,
				VersionPrerelease: VersionPrerelease,
				UI: ui,
			}, nil
		},
	}
}


