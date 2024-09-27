package handler

import (
	"fmt"
	"net/http"

	"belajar-golang-gin-gorm/book"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type BooksHandler struct {
	bookService book.Service
}

func NewBookHandler(bookService book.Service) *BooksHandler {
	return &BooksHandler{bookService}
}

func (h *BooksHandler) GetBooks(c *gin.Context){
	books, err := h.bookService.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	c.JSON(http.StatusOK,gin.H{
		"data": books,
	})
	
}

func (h *BooksHandler) PostBooksHandler(c *gin.Context) {
	var bookRequest book.BookRequest

	// Use ShouldBindJSON instead of ShouldBindBodyWithJSON
	err := c.ShouldBindJSON(&bookRequest)

	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error in field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})

		return
	}

	book, err := h.bookService.Create(bookRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})

	}

	c.JSON(http.StatusOK, gin.H{
		"data":     book,
	})
}
