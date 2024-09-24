package main

import (
	"belajar-golang-gin-gorm/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	v1 := router.Group("/v1")

	v1.GET("/", handler.RootHandler)
	v1.GET("/hello", handler.HelloHandlrer)
	v1.GET("/books/:id/:title", handler.BooksHandler)
	v1.GET("/query", handler.QueryHandler)
	v1.GET("/price", handler.PriceHandler)

	router.POST("/books", handler.PostBooksHandler)

	router.Run()
}




