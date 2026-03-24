package main

import (
	"github.com/youruser/dexter-transport/internal/server"
	_ "github.com/youruser/dexter-transport/docs" // Import swagger docs
)

// @title Dexter Transport API
// @version 1.0
// @description This is a sample server for Dexter Transport.
// @host localhost:8080
// @BasePath /
func main() {
	server.NewServer().Run()
}
