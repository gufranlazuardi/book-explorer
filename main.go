package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main()  {
	router:= gin.Default()

	router.GET("/", rootHandler)
	router.GET("/hello", helloHandlrer)
	router.GET("/books/:id/:title", booksHandler)
	router.GET("/query", queryHandler)
	router.GET("/price", priceHandler)

	router.POST("/books", postBooksHandler)

	router.Run()
}

func rootHandler(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"name" : "Gufran Lazuardi",
		"title" : "Principal Software Engineer",
	})
}

func helloHandlrer(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name" : "Hello",
		"title" : "Selamat pagi",
	})
}

func booksHandler(c *gin.Context){
	id := c.Param("id")
	title := c.Param("title")

	c.JSON(http.StatusOK, gin.H{"id": id , "title": title})
}

func queryHandler(c *gin.Context){
	title := c.Query("title")

	c.JSON(http.StatusOK, gin.H{"title" : title})

}

func priceHandler(c *gin.Context) {
	price := c.Query("price")
	title := c.Query("title")

	c.JSON(http.StatusOK, gin.H{"price" : price, "amount" : title})

	// contoh http://localhost:8080/price?price=20000&amount=10
}

type BookInput struct {
	Title string `json:"title" binding:"required"`
	Price int `json:"price binding:"required,number"`
	SubTitle string `json:"sub_title"`
}

func postBooksHandler(c *gin.Context){
	var bookInput BookInput

	err := c.ShouldBindBodyWithJSON(&bookInput)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		fmt.Println(err)
		return 
	}

	c.JSON(http.StatusOK, gin.H{
		"title" : bookInput.Title,
		"price" : bookInput.Price,
		"sub_title" : bookInput.SubTitle,
	})
}