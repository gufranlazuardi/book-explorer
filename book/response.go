package book

type BookResponse struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Subtitle string `json:"sub_title"`
	Description string `json:"description"`
	Price int `json:"price"`
	Publisher string `json:"publisher"`
	Rating int `json:"rating"`
}