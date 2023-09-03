package main

import (
	"os"
	"todo_api/server"
)

func main() {
	e := server.NewServer()
	e.Start(getPort())
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}
	return ":" + port
}
