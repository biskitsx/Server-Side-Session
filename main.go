package main

import (
	"log"

	"github.com/biskitsx/Server-Side-Session/app"
)

func main() {
	if err := app.SetupAndRunApp(); err != nil {
		log.Fatal(err)
	}
}
