package main

import (
	"log"
	"magnax.ca/add-ca-certificates/pkg/cli"
	"os"
)

func main() {
	app := cli.NewApp()

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
