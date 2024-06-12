package main

import "github.com/Tirivashe/go-fiber-jwt/config"

func main() {
	server := config.NewAPIServer(":8080")
	server.Start()
}
