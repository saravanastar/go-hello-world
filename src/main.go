package main

import (
	"github.com/saravanastar/go-hello-world/src/database"
	"github.com/saravanastar/go-hello-world/src/server"
)

func main() {
	db, _ := database.NewDatabaseClient()
	server := server.NewEchoServer(db)
	server.Start()
}
