package main

import (
	"example.com/myapp/db"
	"example.com/myapp/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisterEventRoutes(server)
	server.Run(":8080")
}
