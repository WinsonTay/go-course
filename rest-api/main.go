package main

import (
	"example.com/eventbooking-rest-api/db"
	"example.com/eventbooking-rest-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8080") // localhost
}
