package main

import (
	"mag/controllers"
	_ "mag/docs"
	"mag/models"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	app := gin.Default()

	models.ConnectDatabase()

	v1 := app.Group("/api/v1")
	{
		v1.GET("/books", controllers.FindBooks)
		v1.GET("/books/:id", controllers.FindBook)
		v1.POST("/books", controllers.CreateBook)
		v1.PATCH("/books/:id", controllers.UpdateBook)
		v1.DELETE("/books/:id", controllers.DeleteBook)
		app.GET("/swagger/v1/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	app.Run()

}
