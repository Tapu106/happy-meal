package main

import (
	"github.com/fconnect/happy-meal/internal/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080")

}
