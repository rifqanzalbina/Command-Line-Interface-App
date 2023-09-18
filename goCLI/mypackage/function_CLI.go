package mypackage

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func StartedCLI() {
	app := &cli.App{
		Name:  "boom",
		Usage: "make an explosive entrance",
		Action: func(*cli.Context) error {
			fmt.Println("boom! I say!")
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func GreetCLI() {
	app := &cli.App{
		Name:  "greet",
		Usage: "Fight the loneliness",
		Action: func(*cli.Context) error {
			fmt.Println("Hello Friend !")
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func ArgumentCLI() {
	app := &cli.App{
		Action: func(cCtx *cli.Context) error {
			fmt.Printf("Hello Everyone %q", cCtx.Args().Get(0))
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}


// FlagCLI 1
func FlagCLI() {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "lang",
				Value: "english",
				Usage: "language for the greeting",
			},
		},
		Action: func(cCtx *cli.Context) error {
			name := "Nefertiti"
			if cCtx.NArg() > 0 {
				name = cCtx.Args().Get(0)
			}
			if cCtx.String("lang") == "spanish" {
				fmt.Println("Hola", name)
			} else {
				fmt.Println("Hello", name)
			}
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}


// FlagCLI2
func FlagCLI2() {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "config",
				Aliases: []string{"c"},
				Usage:   "Load configuration from 'FILE'",
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}

}

