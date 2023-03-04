package main

import (
	"fmt"
	"log"
	"os"

	"github.com/mtstnt/mog/commands"
	"github.com/mtstnt/mog/config"
	"github.com/urfave/cli/v2"
)

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}

func run() error {
	if err := config.LoadGlobals(); err != nil {
		return err
	}
	app := &cli.App{
		Name:  "mog",
		Usage: "Generate and do database migrations with simple tooling.",
		Action: func(ctx *cli.Context) error {
			fmt.Println("Saying hello!")
			return nil
		},
		Commands: []*cli.Command{
			commands.InitCommand(),
			commands.MigrateCommand(),
			commands.NewCommand(),
		},
	}

	return app.Run(os.Args)
}
