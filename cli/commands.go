package cli

import (
	"os"

	"github.com/mitchellh/cli"
	"github.com/patdhlk/cli/command"
)

func Commands(metaPtr *command.Meta) map[string]cli.CommandFactory {
	if metaPtr == nil {
		metaPtr = new(command.Meta)
	}

	meta := *metaPtr
	if meta.Ui == nil {
		meta.Ui = &cli.BasicUi{
			Writer:      os.Stdout,
			ErrorWriter: os.Stderr,
		}
	}

	return map[string]cli.CommandFactory{
		"server": func() (cli.Command, error) {
			return &command.ServerCommand{
				Ui: meta.Ui,
			}, nil
		},
		"version": func() (cli.Command, error) {
			ver := Version
			rel := VersionPrerelease

			return &command.VersionCommand{
				Revision:          GitCommit,
				Version:           ver,
				VersionPrerelease: rel,
				Ui:                meta.Ui,
			}, nil
		},
	}
}
