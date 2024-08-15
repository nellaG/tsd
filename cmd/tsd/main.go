package main

import (
	"log"
	"os"

	"github.com/nellaG/tsd/internal/cli"
)

func main() {
	app := cli.NewApp()
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
