package main

import (
	"belajar-golang-gin-gorm/book"
	connection "belajar-golang-gin-gorm/config"
	"belajar-golang-gin-gorm/handler"
	"belajar-golang-gin-gorm/router"
)

func main() {

	db := connection.Connect()

	db.AutoMigrate(&book.Book{})

	bookRepository := book.NewRepository(db)
	bookService := book.NewService(bookRepository)
	bookHandler := handler.NewBookHandler(bookService)

	r := router.SetupRouter(bookHandler)
	
	r.Run()
}

// main
// handler
// service
// repository
// db
// postgres
