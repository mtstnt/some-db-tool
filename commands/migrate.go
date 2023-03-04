package commands

import (
	"github.com/mtstnt/mog/config"
	"github.com/urfave/cli/v2"
)

func MigrateCommand(conf *config.GlobalInfo) *cli.Command {
	cmd := &cli.Command{
		Name:  "migrate",
		Usage: "Migrates your current migrations to database.",
		Flags: migrateFlags(),
		Action: func(ctx *cli.Context) error {
			return migrateAction(ctx, conf)
		},
	}
	return cmd
}

func migrateFlags() []cli.Flag {
	return []cli.Flag{}
}

func migrateAction(ctx *cli.Context, conf *config.GlobalInfo) error {
	return nil
}
