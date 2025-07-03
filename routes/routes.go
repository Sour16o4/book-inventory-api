package routes

import (
	"book-inventory-api/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	{
		v1.GET("/books", controllers.GetBooks)
		v1.GET("/books/:id", controllers.GetBookByID)
		v1.POST("/books", controllers.CreateBook)
		v1.PUT("/books/:id", controllers.UpdateBook)
		v1.DELETE("/books/:id", controllers.DeleteBook)
	}
}
