package cli

import (
	"fmt"

	"github.com/nellaG/tsd/internal/build"
	"github.com/nellaG/tsd/internal/converter"
	"github.com/urfave/cli/v2"
)

func NewApp() *cli.App {
	app := &cli.App{
		Name:    "tsd",
		Usage:   "Easily convert UNIX timestamp to date ISO8601 format",
		Version: build.Version,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "test",
				Aliases: []string{"t"},
				Usage:   "dummy test flag",
			},
		},
		Action: func(c *cli.Context) error {
			if c.NArg() < 1 {
				return fmt.Errorf("no timestamp. please provide: ")
			}

			timestamp := c.Args().First()
			date, err := converter.TimestampToDate(timestamp)
			if err != nil {
				return fmt.Errorf("invalid value: %v", err)
			}
			fmt.Println(date)
			return nil
		},
	}
	return app
}
