package main

import (
	"log"

	"github.com/mariolazzari/micro-go/internal/database"
	"github.com/mariolazzari/micro-go/internal/server"
)

func main() {
	db, err := database.NewDatabaseClient()
	if err != nil {
		log.Fatalf("failed to initialize Database Client: %s", err)
	}
	srv := server.NewEchoServer(db)
	if err := srv.Start(); err != nil {
		log.Fatal(err.Error())
	}
}
