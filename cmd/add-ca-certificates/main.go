package main

import (
	"log"
	"os"

	"magnax.ca/add-ca-certificates/pkg/cli"
)

func main() {
	app := cli.NewApp()

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
