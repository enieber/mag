package controllers

import (
	"mag/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Buy Product
// @Summary Buy Product
// @Schemes
// @Description Buy Product
// @Tags products
// @Accept json
// @Param buyproduct body models.SalesInput true "SalesInput to Buy"
// @Success 200 {object} models.Product
// @Router /api/v1/sales/buy [post]
func BuyProduct(ctx *gin.Context) {
	var input models.SalesInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User

	if err := models.DB.Where("id = ?", input.IdUser).First(&user).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "user not found!"})
		return
	}

	var product models.Product

	if err := models.DB.Where("id = ?", input.IdProduct).First(&product).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "product not found!"})
		return
	}

	sale := models.Sale{ProductID: product.ID, UserID: user.ID, Status: "Pending"}
	models.DB.Create(&sale)
	transaction := models.Transaction{SaleID: sale.ID, Status: sale.Status}
	models.DB.Create(&transaction)
	transactionReturn := models.TransactionReturn{ID: transaction.ID, Status: transaction.Status}
	ctx.JSON(http.StatusOK, gin.H{"data": transactionReturn})

}
