package main

import (
	"todo_api/server"
)

func main() {
	e := server.NewServer()
	e.Start(":5000")
}
