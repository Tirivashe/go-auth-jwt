package main

import (
	"log"

	"github.com/Tirivashe/go-fiber-jwt/config"
	"github.com/Tirivashe/go-fiber-jwt/db"
)

func init() {
	db.ConnectToDB()
}

func main() {
	server := config.NewAPIServer(":8080")
	defer db.DB.Close()
	log.Fatal(server.Start())
}
