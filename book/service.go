package book

import "fmt"

type Service interface {
	FindAll() ([]Book, error)
	FindById(ID int) (Book, error)
	Create(book BookRequest) (Book, error)
	Update(ID int, book BookRequest) (Book, error)
	Delete(ID int) (Book, error)
}

type service struct {
	repository Repository
}

func NewService (repository Repository) *service {
	return &service{repository}
}

func(s service) FindAll()([]Book, error) {
	books, err := s.repository.FindAll()
	return books, err
}

func(s *service) FindById(ID int) (Book, error) {
	books, err := s.repository.FindById(ID)
	return books, err
}

func(s service) Create(bookRequest BookRequest)(Book, error) {
	price, _ := bookRequest.Price.Int64()
	rating, _ := bookRequest.Rating.Int64()

	book := Book{
		Title: bookRequest.Title,
		Price: int(price),
		Description: bookRequest.Description,
		Publisher: bookRequest.Publisher,
		SubTitle: bookRequest.SubTitle,
		Rating: float32(rating),
	}

	newBook, err := s.repository.Create(book)
	return newBook, err
}

func(s service) Update(ID int, bookRequest BookRequest)(Book, error) {
	// cari dulu data bukunya
	book, err := s.repository.FindById(ID)
	if err != nil {
		fmt.Println("Error", err)
	}

	// baru update
	price, _ := bookRequest.Price.Int64()
	rating, _ := bookRequest.Rating.Int64()

	book.Title = bookRequest.Title
	book.Price = int(price)
	book.Description = bookRequest.Description
	book.Publisher = bookRequest.Publisher
	book.SubTitle = bookRequest.SubTitle
	book.Rating = float32(rating)

	newBook, err := s.repository.Update(book)
	return newBook, err
}

func(s service) Delete(ID int)(Book, error) {
	// cari dulu data bukunya
	book, err := s.repository.FindById(ID)
	if err != nil {
		fmt.Print("Error", err)
	}

	// baru update
	newBook, err := s.repository.Delete(book)
	return newBook, err
}