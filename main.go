package main

import (
	"belajar-golang-gin-gorm/book"
	"belajar-golang-gin-gorm/handler"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	dsn := "host=localhost user=macbook password=1234 dbname=belajar-golang-gin-gorm port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Db connection error")
	}

	fmt.Println("Database connection succeed")

	db.AutoMigrate(&book.Book{})

	bookRepository := book.NewRepository(db)
	bookService := book.NewService(bookRepository)

	bookRequest := book.BookRequest{
		Title: "Buku Keren",
		Price: "200000",
	}

	bookService.Create(bookRequest)
		
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
