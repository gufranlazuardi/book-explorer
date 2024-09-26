package main

import (
	"belajar-golang-gin-gorm/book"
	"belajar-golang-gin-gorm/handler"
	"belajar-golang-gin-gorm/router"
	"fmt"
	"log"

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

	r := router.SetupRouter(bookHandler)
	
	r.Run()
}

// main
// handler
// service
// repository
// db
// postgres
