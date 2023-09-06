package main

import (
	"log"

	"github.com/b0gochort/microservices/internal/app"
)

func main() {
	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
