package main

import (
	"os"
	"comentarismo-spam/server"
)

var Port = os.Getenv("PORT")

func main() {
	if Port == "" {
		Port = "3000"
	}
	server.StartServer(Port)
}
