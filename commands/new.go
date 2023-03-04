package commands

import "github.com/urfave/cli/v2"

func NewCommand() *cli.Command {
	return &cli.Command{
		Name:   "new",
		Usage:  "Generates a new migration file.",
		Action: newAction,
	}
}

func newAction(ctx *cli.Context) error {
	return nil
}
