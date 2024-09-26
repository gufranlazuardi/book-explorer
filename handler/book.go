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

func (handler *BooksHandler) RootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name":  "Gufran Lazuardi",
		"title": "Principal Software Engineer",
	})
}

func  (h *BooksHandler) HelloHandlrer(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name":  "Hello",
		"title": "Selamat pagi",
	})
}

func  (h *BooksHandler) BooksHandler(c *gin.Context) {
	id := c.Param("id")
	title := c.Param("title")

	c.JSON(http.StatusOK, gin.H{"id": id, "title": title})
}

func  (h *BooksHandler) PriceHandler(c *gin.Context) {
	price := c.Query("price")
	title := c.Query("title")

	c.JSON(http.StatusOK, gin.H{"price": price, "amount": title})
	// contoh http://localhost:8080/price?price=20000&amount=10
}

func  (h *BooksHandler) QueryHandler(c *gin.Context) {
	title := c.Query("title")

	c.JSON(http.StatusOK, gin.H{"title": title})
}

func  (h *BooksHandler) PostBooksHandler(c *gin.Context) {
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
