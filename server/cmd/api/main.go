package main

import (
	"log"

	"vow/server/internal/app"
)

func main() {
	api, err := app.New()
	if err != nil {
		log.Fatal(err)
	}

	if err := api.Run(); err != nil {
		log.Fatal(err)
	}
}
