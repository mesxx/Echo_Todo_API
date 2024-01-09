package main

import (
	"echo_todo_api/server"
	"os"
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
