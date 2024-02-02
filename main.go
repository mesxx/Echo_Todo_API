package main

import (
	"os"

	"github.com/mesxx/Echo_Todo_API/server"
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
