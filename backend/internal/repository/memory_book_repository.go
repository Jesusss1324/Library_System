package repository

import "library-system/internal/model"

type MemoryBookRepository struct {
	books  []model.Book
	nextID int
}

func NewMemoryBookRepository() *MemoryBookRepository {
	return &MemoryBookRepository{
		books:  []model.Book{},
		nextID: 1,
	}
}

func (r *MemoryBookRepository) CreateBook(book model.Book) (model.Book, error) {
	book.ID = r.nextID
	r.nextID++

	r.books = append(r.books, book)

	return book, nil
}

func (r *MemoryBookRepository) GetBooks() ([]model.Book, error) {
	return r.books, nil
}
