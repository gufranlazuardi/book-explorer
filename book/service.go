package book

type Service interface {
	FindAll() ([]Book, error)
	FindById(ID int) (Book, error)
	Create(book BookRequest) (Book, error)
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

func(s service) FindById(ID int)(Book, error) {
	books, err := s.repository.FindById(ID)
	return books, err
}

func(s service) Create(bookRequest BookRequest)(Book, error) {

	price, _ := bookRequest.Price.Int64()

	book := Book{
		Title: bookRequest.Title,
		Price: int(price),
	}

	newBook, err := s.repository.Create(book)
	return newBook, err
}
