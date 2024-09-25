package book

import "time"

type Book struct {
	ID int
	Title string
	Description string
	Price int
	Rating float32
	Publisher string
	CreatedAt time.Time
	UpdatedAt time.Time
}


