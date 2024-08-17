package main

import (
	route "drink-api/cmd/api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	route.DrinksRoutes(server)

	server.Run(":8080")
}
