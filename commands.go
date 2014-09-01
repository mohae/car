// Initializes the Commands struct for the application.
// New commands need to be added to the CommandFactory map.
package main

import (
	"os"

	"github.com/mitchellh/cli"
	"github.com/mohae/clitemplate/command"
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
	}
}


