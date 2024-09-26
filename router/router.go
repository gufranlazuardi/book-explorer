package router

import (
	"belajar-golang-gin-gorm/handler"

	"github.com/gin-gonic/gin"
)

func SetupRouter(bookHandler *handler.BooksHandler) *gin.Engine {
	router := gin.Default()
	v1 := router.Group("v1")
	{
		v1.GET("/", bookHandler.RootHandler)
		v1.GET("/hello", bookHandler.HelloHandlrer)
		v1.GET("/books/:id/:title", bookHandler.BooksHandler)
		v1.GET("/query", bookHandler.QueryHandler)
		v1.GET("/price", bookHandler.PriceHandler)
		v1.POST("/books", bookHandler.PostBooksHandler)
	}

	return router
}