package controllers

import (
	"mag/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// Create User
// @Summary Create User
// @Schemes
// @Description Create User
// @Tags users
// @Accept json
// @Param createBook body models.UserInput true "UserInput to create"
// @Success 200 {object} models.UserReturn
// @Router /api/v1/users [post]
func CreateUser(ctx *gin.Context) {
	var input models.UserInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user := models.User{Name: input.Name, Email: input.Email}
	models.DB.Create(&user)

	ctx.JSON(http.StatusOK, gin.H{"data": user})

}
