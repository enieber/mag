package controllers

import (
	"mag/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// List Books
// @Summary List all books
// @Schemes
// @Description List books from database
// @Tags books
// @Accept json
// @Produce json
// @Success 200 {array} models.Book
// @Router /api/v1/books [get]
func FindBooks(ctx *gin.Context) {
	var books []models.Book
	models.DB.Find(&books)
	if books != nil {
		ctx.JSON(http.StatusOK, gin.H{"data": books})
	} else {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "books not found"})
	}
}

// @BasePath /api/v1

// Create book
// @Summary Create book
// @Schemes
// @Description Create book
// @Tags books
// @Accept json
// @Param createBook body models.CreateBookInput true "CreateBookInput to create"
// @Success 200 {object} models.Book
// @Router /api/v1/books [post]
func CreateBook(ctx *gin.Context) {
	var input models.CreateBookInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	book := models.Book{Title: input.Title, Author: input.Author}
	models.DB.Create(&book)

	ctx.JSON(http.StatusOK, gin.H{"data": book})

}

// @BasePath /api/v1

// Find Book by id
// @Summary find all books
// @Schemes
// @Description books from database
// @Tags books
// @Accept json
// @Param id path int true "id of book"
// @Success 200 {object} models.Book
// @Router /api/v1/books/{id} [get]
func FindBook(ctx *gin.Context) { // Get model if exist
	var book models.Book

	if err := models.DB.Where("id = ?", ctx.Param("id")).First(&book).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": book})
}

// @BasePath /api/v1

// Update book
// @Summary Update book
// @Schemes
// @Description Update book
// @Tags books
// @Accept json
// @Param id path int true "id of book"
// @Param UpdateBookInput body models.UpdateBookInput true "UpdateBookInput to create"
// @Success 200 {object} models.Book
// @Router /api/v1/books/{id} [patch]
func UpdateBook(ctx *gin.Context) {
	// Get model if exist
	var book models.Book
	if err := models.DB.Where("id = ?", ctx.Param("id")).First(&book).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input models.UpdateBookInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&book).Updates(input)

	ctx.JSON(http.StatusOK, gin.H{"data": book})
}

// @BasePath /api/v1

// Deleted book
// @Summary Deleted book
// @Schemes
// @Description Deleted book
// @Tags books
// @Accept json
// @Param id path int true "id of book"
// @Success 200 {string} data
// @Router /api/v1/books/{id} [delete]
func DeleteBook(ctx *gin.Context) {
	// Get model if exist
	var book models.Book
	if err := models.DB.Where("id = ?", ctx.Param("id")).First(&book).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&book)

	ctx.JSON(http.StatusOK, gin.H{"data": true})
}
