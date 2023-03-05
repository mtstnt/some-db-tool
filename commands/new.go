package commands

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/gobeam/stringy"
	"github.com/mtstnt/mog/config"
	"github.com/urfave/cli/v2"
)

func NewCommand(conf *config.GlobalInfo) *cli.Command {
	return &cli.Command{
		Name:  "new",
		Usage: "Generates a new migration file.",
		Action: func(ctx *cli.Context) error {
			return newAction(ctx, conf)
		},
	}
}

func newAction(ctx *cli.Context, conf *config.GlobalInfo) error {
	args := ctx.Args()
	migrationName := args.First()

	if migrationName == "" {
		return fmt.Errorf("got blank string, expected migration name")
	}

	path := ctx.String("path")
	if path == "" {
		path = filepath.Join(conf.CurrentPath, "mog.config.yml")
	}

	mogConfig, err := config.ReadConfig(path)
	if err != nil {
		return fmt.Errorf("newAction: error loading mogconfig: %w", err)
	}

	migrationDir := mogConfig.Project.MigrationDir
	if migrationDir == "" {
		migrationDir = filepath.Join(conf.CurrentPath, "migrations")
	}

	currentTime := time.Now()

	migrationProcessedFilename := fmt.Sprintf(
		"%d_%02d%02d%d_%s.yml",
		currentTime.Unix(),
		currentTime.Day(),
		currentTime.Month(),
		currentTime.Year(),
		stringy.New(migrationName).SnakeCase().ToLower(),
	)

	if err := os.MkdirAll(migrationDir, os.ModeDir); err != nil {
		return fmt.Errorf("newAction: err mkdir migrations dir: %w", err)
	}

	{
		migrationFilename := filepath.Join(conf.CurrentPath, migrationDir, migrationProcessedFilename)
		fptr, err := os.Create(migrationFilename)
		if err != nil {
			return fmt.Errorf("newAction: error creating file %s: %w", migrationFilename, err)
		}
		defer fptr.Close()

		templateFilename := filepath.Join(conf.ExecutablePath, "templates", "migration.yml")
		fptr2, err := os.Open(templateFilename)
		if err != nil {
			return fmt.Errorf("newAction: error loading file %s: %w", templateFilename, err)
		}
		defer fptr2.Close()

		if _, err := io.Copy(fptr, fptr2); err != nil {
			return fmt.Errorf("newAction: error copying from templates to src: %w", err)
		}
	}

	return nil
}
