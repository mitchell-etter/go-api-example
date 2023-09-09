package main

import (
	"log"

	"go-api-example/internal/apiServer"
)

func main() {
	log.Println("Starting up")

	server := apiServer.NewApiServer()
	server.Run()

	log.Println("Exiting")
}
