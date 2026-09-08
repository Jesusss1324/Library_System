package repository

import (
	"context"
	"database/sql"
	"fmt"

	"library-system/internal/model"
)

type BookRepository struct {
	db *sql.DB
}

func NewBookRepository(db *sql.DB) *BookRepository {
	return &BookRepository{
		db: db,
	}
}

func (r *BookRepository) CreateBook(book model.Book) (model.Book, error) {
	query := `
		INSERT INTO books (title, author)
		OUTPUT INSERTED.id
		VALUES (@p1, @p2);
	`

	err := r.db.QueryRowContext(
		context.Background(),
		query,
		book.Title,
		book.Author,
	).Scan(&book.ID)

	if err != nil {
		return model.Book{}, fmt.Errorf("failed to create book: %w", err)
	}

	return book, nil
}

func (r *BookRepository) GetBooks() ([]model.Book, error) {
	query := `
		SELECT id, title, author
		FROM books
		ORDER BY id;
	`

	rows, err := r.db.QueryContext(context.Background(), query)
	if err != nil {
		return nil, fmt.Errorf("failed to get books: %w", err)
	}
	defer rows.Close()

	books := []model.Book{}

	for rows.Next() {
		var book model.Book

		if err := rows.Scan(
			&book.ID,
			&book.Title,
			&book.Author,
		); err != nil {
			return nil, fmt.Errorf("failed to scan book: %w", err)
		}

		books = append(books, book)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error reading books: %w", err)
	}

	return books, nil
}
