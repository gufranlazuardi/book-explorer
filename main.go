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
	bookHandler := handler.NewBookHandler(bookService)


	// bookRequest := book.BookRequest{
	// 	Title: "Buku Keren",
	// 	Price: "200000",
	// }

	router := gin.Default()

	v1 := router.Group("/v1")

	v1.GET("/", bookHandler.RootHandler)
	v1.GET("/hello", bookHandler.HelloHandlrer)
	v1.GET("/books/:id/:title", bookHandler.BooksHandler)
	v1.GET("/query", bookHandler.QueryHandler)
	v1.GET("/price", bookHandler.PriceHandler)

	v1.POST("/books", bookHandler.PostBooksHandler)

	router.Run()
}

// main
// handler
// service
// repository
// db
// postgres
