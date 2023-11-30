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
		v1.POST("/products", controllers.CreateProduct)
		v1.POST("/users", controllers.CreateUser)
		v1.POST("/sales/buy", controllers.BuyProduct)
		v1.PUT("/sales/payment", controllers.UpdateTransaction)
		app.GET("/swagger/v1/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	app.Run()

}
