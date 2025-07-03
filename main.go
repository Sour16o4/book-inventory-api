package main

import (
	"book-inventory-api/database"
	"book-inventory-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()

	r := gin.Default()
	routes.SetupRoutes(r)

	// Start server
	r.Run(":8080") // 🚀 This line is CRITICAL
}
