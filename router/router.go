package router

import (
	"belajar-golang-gin-gorm/handler"

	"github.com/gin-gonic/gin"
)

func SetupRouter(bookHandler *handler.BooksHandler) *gin.Engine {
	router := gin.Default()
	v1 := router.Group("v1")
	{
		v1.GET("/books", bookHandler.GetBooks)
		v1.GET("books/:id", bookHandler.GetBook)
		v1.POST("/books", bookHandler.CreateBook)
		v1.PUT("/books/:id",bookHandler.UpdateBook)
	}

	return router
}