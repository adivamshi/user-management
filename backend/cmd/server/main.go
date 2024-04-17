package main

import (
	"log"
	"os"
	"user-management-api/server"
)

func main() {
	defaultPort := "8081"
	port := os.Getenv("PORT")
	if port == "" {
		log.Printf("PORT environment variable not set, defaulting to %s", defaultPort)
		port = defaultPort
	}
	server.Run(port)
}
