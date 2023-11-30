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
// @Tags sales
// @Accept json
// @Param buyproduct body models.SalesInput true "SalesInput to Buy"
// @Success 200 {object} models.TransactionReturn
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

// Update Payment
// @Summary Update Payment
// @Schemes
// @Description update status of transaction when payment updated
// @Tags sales
// @Accept json
// @Param transaction body models.TransactionInput true "TransactionInput to update paymnentOk"
// @Success 200 {object} models.TransactionReturn
// @Router /api/v1/sales/payment [put]
func UpdateTransaction(ctx *gin.Context) {
	var input models.TransactionInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.Status == "Pending" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Status not Allow"})
		return
	}

	if input.Status != "PaymentOk" ||
		input.Status != "PaymentCancell" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Status not Allow"})
		return
	}

	var transaction models.Transaction

	if err := models.DB.Where("id = ?", input.ID).First(&transaction).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "transaction not found!"})
		return
	}

	var sale models.Sale

	if err := models.DB.Where("id = ?", transaction.SaleID).First(&sale).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "sale not found!"})
		return
	}

	sale.Status = input.Status
	models.DB.Save(&sale)
	transaction.Status = input.Status
	models.DB.Save(&transaction)
	transactionReturn := models.TransactionReturn{ID: transaction.ID, Status: transaction.Status}
	ctx.JSON(http.StatusOK, gin.H{"data": transactionReturn})
}
