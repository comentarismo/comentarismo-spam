package main

import (
	"comentarismo-spam/server"
	"os"
)

var Port = os.Getenv("PORT")

func main() {
	if Port == "" {
		Port = "3004"
	}
	server.StartServer(Port)
}
