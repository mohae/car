package main

import (
	"os"

	"github.com/mitchellh/cli"
	"github.com/mohae/baller/command"
)

// Commands is the mapping of all available baller commands.
var Commands map[string]cli.CommandFactory

func init() {
	ui := &cli.BasicUi{Writer: os.Stdout}
	Commands = map[string]cli.CommandFactory{
		"create": func() (cli.Command, error) {
			return &command.CreateCommand{
				UI: ui,
			}, nil
		},
		"delete": func() (cli.Command, error) {
			return &command.DeleteCommand{
				UI: ui,
			}, nil
		},
		"version": func() (cli.Command, error) {
			return &command.VersionCommand{
				Revision: GitCommit,
				Version: Version,
				VersionPrerelease: VersionPrerelease,
				UI: ui,
			}, nil
		},
	}
}

