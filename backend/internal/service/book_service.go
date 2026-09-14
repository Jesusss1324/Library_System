package service

import "library-system/internal/model"

type BookRepository interface {
	CreateBook(book model.Book) (model.Book, error)
	GetBooks() ([]model.Book, error)
}

type BookService struct {
	repository BookRepository
}

func NewBookService(repository BookRepository) *BookService {
	return &BookService{
		repository: repository,
	}
}

func (s *BookService) CreateBook(book model.Book) (model.Book, error) {
	return s.repository.CreateBook(book)
}

func (s *BookService) GetBooks() ([]model.Book, error) {
	return s.repository.GetBooks()
}
