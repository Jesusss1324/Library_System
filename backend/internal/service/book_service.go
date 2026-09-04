package service

import "library-system/internal/model"

var books = []model.Book{}

func CreateBook(book model.Book) model.Book {
	book.ID = len(books) + 1
	books = append(books, book)

	return book
}

func GetBooks() []model.Book {
	return books
}
