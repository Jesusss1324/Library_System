package service

import "library-system/internal/model"

type BookService struct {
	books []model.Book
}

func NewBookService() *BookService {
	return &BookService{
		books: []model.Book{},
	}
}

func (s *BookService) CreateBook(book model.Book) model.Book {
	book.ID = len(s.books) + 1

	s.books = append(s.books, book)

	return book
}

func (s *BookService) GetBooks() []model.Book {
	return s.books
}
