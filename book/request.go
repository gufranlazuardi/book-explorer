package book

import (
	"encoding/json"
)

type BookRequest struct {
	Title    string      `json:"title" binding:"required"`
	SubTitle string      `json:"sub_title"`
	Price    json.Number `json:"price" binding:"required,number"`
	Description string 	`json:"description" binding:"required"`
	Publisher string `json:"publisher" binding:"required"`
	Rating json.Number `json:"rating" binding:"required"`
}
