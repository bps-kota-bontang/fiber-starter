package main

import (
	"fmt"

	"log"

	"github.com/bps-kota-bontang/fiber-starter/di"
)

func main() {
	app, err := di.InitializeApp()
	if err != nil {
		log.Fatalf("Failed to initialize app: %v", err)
	}

	log.Fatal(app.App.Listen(fmt.Sprintf(":%s", app.Config.AppPort)))
}
