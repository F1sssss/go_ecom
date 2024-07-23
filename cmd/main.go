package main

import (
	"log"

	"github.com/F1sssss/go_ecom/cmd/api"
	"github.com/F1sssss/go_ecom/db"
)

func main() {

	db, err := db.NewDB()

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	server := api.NewServer()
	server.Start()

}
