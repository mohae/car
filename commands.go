package main

import (
	"os"

	"github.com/mohae/car/command"
	"github.com/mohae/cli"
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
		"extract": func() (cli.Command, error) {
			return &command.ExtractCommand{
				UI: ui,
			}, nil
		},
		"version": func() (cli.Command, error) {
			return &command.VersionCommand{
				Revision:          GitCommit,
				Version:           Version,
				VersionPrerelease: VersionPrerelease,
				UI:                ui,
			}, nil
		},
	}
}
