package controllers

import (
	"mag/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Create Product
// @Summary Create Product
// @Schemes
// @Description Create Product
// @Tags products
// @Accept json
// @Param createBook body models.ProductInput true "ProductInput to create"
// @Success 200 {object} models.ProductReturn
// @Router /api/v1/products [post]
func CreateProduct(ctx *gin.Context) {
	var input models.ProductInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	product := models.Product{Name: input.Name, Type: input.Type}
	models.DB.Create(&product)

	ctx.JSON(http.StatusOK, gin.H{"data": product})
}
