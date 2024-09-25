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
	// crud

	// book := book.Book{}
	// book.Title = "PKI"
	// book.Price = 70000
	// book.Publisher = "Partai Komunis"
	// book.Rating = 5.5
	// book.Description = "PKI HARAM"

	// err = db.Create(&book).Error
	// if err != nil {
	// 	fmt.Println("==========================")
	// 	fmt.Println("Error creating book record")
	// 	fmt.Println("==========================")
	// }

	var fieldBook book.Book

	// err = db.Debug().Where("Title = ?", "PKI").First(&fieldBook).Error
	// if err != nil {
	// 	fmt.Println("==========================")
	// 	fmt.Println("Error catching title:", err)
	// 	fmt.Println("==========================")
	// }

	// // Update the title
	// fieldBook.Title = "NEW PKI IS BORN"

	// // Save the updated record
	// err = db.Save(&fieldBook).Error
	// if err != nil {
	// 	fmt.Println("==========================")
	// 	fmt.Println("Error updating title:", err)
	// 	fmt.Println("==========================")
	// } else {
	// 	fmt.Println("Title successfully updated to:", fieldBook.Title)
	// }


	err = db.Debug().Where("id = ?", 1).First(&fieldBook).Error
	if err != nil {
		fmt.Println("==========================")
		fmt.Println("Error updating title:", fieldBook)
		fmt.Println("==========================")
	}

	err = db.Debug().Delete(&fieldBook).Error
	if err != nil {
		fmt.Println("==========================")
		fmt.Println("Success deleted ID :", fieldBook)
		fmt.Println("==========================")
	}

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
