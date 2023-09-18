package mypackage

import (
	"log"
	"os"
	"github.com/urfave/cli/v2"
)

func ExitCodes() {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "ginger-crouton",
				Usage: "is it in the soup",
			},
		},
		Action: func(ctx *cli.Context) error {
			if !ctx.Bool("ginger-crouton") {
				return cli.Exit("Ginger croutons are not in the soup", 86)
			}
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}

}
