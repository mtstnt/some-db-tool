package commands

import (
	"github.com/urfave/cli/v2"
)

func MigrateCommand() *cli.Command {
	cmd := &cli.Command{
		Name:   "migrate",
		Usage:  "Migrates your current migrations to database.",
		Flags:  migrateFlags(),
		Action: migrateAction,
	}
	return cmd
}

func migrateFlags() []cli.Flag {
	return []cli.Flag{}
}

func migrateAction(ctx *cli.Context) error {
	return nil
}
