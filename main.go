package main

import (
	"go-cli/app"
	"log"
	"os"
	// imports as package "cli"
)

func main() {
	app := app.NewCli()
	
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}