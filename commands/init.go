package commands

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/mtstnt/mog/config"
	"github.com/urfave/cli/v2"
)

var (
	templateFilename string = "mog.config.yml"
)

func InitCommand(conf *config.GlobalInfo) *cli.Command {
	cmd := &cli.Command{
		Name:  "init",
		Usage: "Generates the mog.config.yml file and `mog` directory.",
		Flags: initFlags(),
		Action: func(ctx *cli.Context) error {
			return initAction(ctx, conf)
		},
	}
	return cmd
}

func initFlags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:  "path",
			Usage: "The path where the root of the migration project will be located.",
		},
	}
}

func initAction(ctx *cli.Context, conf *config.GlobalInfo) error {
	path := conf.CurrentPath

	pathFlag := ctx.String("path")
	if pathFlag != "" {
		path = conf.CurrentPath + pathFlag
	}

	configDstPath := filepath.Join(path, templateFilename)
	configSrcPath := filepath.Join(conf.ExecutablePath, "templates", templateFilename)

	newFilePtr, err := os.Create(configDstPath)
	if err != nil {
		return fmt.Errorf("error creating config file: %w", err)
	}
	defer newFilePtr.Close()

	templatePtr, err := os.Open(configSrcPath)
	if err != nil {
		return fmt.Errorf("error opening template file %s: %w", configSrcPath, err)
	}
	defer templatePtr.Close()

	_, err = io.Copy(newFilePtr, templatePtr)
	if err != nil {
		return fmt.Errorf("error copying contents of template file to %s: %w", configDstPath, err)
	}

	if err := newFilePtr.Sync(); err != nil {
		return fmt.Errorf("error syncing new config file: %w", err)
	}

	fmt.Printf("Config file: %s has been created successfully.\n", configDstPath)
	return nil
}
