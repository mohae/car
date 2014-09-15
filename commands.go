package main

import (
	"os"

<<<<<<< HEAD
	"github.com/mohae/car/command"
	"github.com/mohae/cli"
=======
	"github.com/mitchellh/cli"
	"github.com/mohae/car/command"
>>>>>>> 2e926bf1982fb46d62edad497bf67f5e2bff12e7
)

// Commands is the mapping of all available car commands.
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
<<<<<<< HEAD
				Revision:          GitCommit,
				Version:           Version,
				VersionPrerelease: VersionPrerelease,
				UI:                ui,
=======
				Revision: GitCommit,
				Version: Version,
				VersionPrerelease: VersionPrerelease,
				UI: ui,
>>>>>>> 2e926bf1982fb46d62edad497bf67f5e2bff12e7
			}, nil
		},
	}
}
<<<<<<< HEAD
=======

>>>>>>> 2e926bf1982fb46d62edad497bf67f5e2bff12e7
