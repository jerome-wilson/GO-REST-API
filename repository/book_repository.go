package repository

import (
	"database/sql"
	"errors"

	"github.com/jerome-wilson/GO-REST-API/models"
)

// BookRepository handles DB operations for books.
type BookRepository struct {
	DB *sql.DB
}

// NewBookRepository creates a new BookRepository with a DB connection.
// NewBookRepository creates a new BookRepository using the provided DB.
// The DB should be created/managed by the caller (allows easier testing).
func NewBookRepository(database *sql.DB) *BookRepository {
	return &BookRepository{DB: database}
}

// InsertBook inserts a new book into the database.
func (r *BookRepository) InsertBook(book *models.Book) (int64, error) {
	query := "INSERT INTO books (title, author, published_year) VALUES (?, ?, ?)"
	result, err := r.DB.Exec(query, book.Title, book.Author, book.Year)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

// GetBookByID retrieves a book by its ID.
func (r *BookRepository) GetBookByID(id int64) (*models.Book, error) {
	query := "SELECT id, title, author, published_year FROM books WHERE id = ?"
	row := r.DB.QueryRow(query, id)
	book := &models.Book{}
	err := row.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
	if err == sql.ErrNoRows {
		return nil, errors.New("book not found")
	}
	if err != nil {
		return nil, err
	}
	return book, nil
}

// UpdateBook updates an existing book.
func (r *BookRepository) UpdateBook(book *models.Book) error {
	query := "UPDATE books SET title = ?, author = ?, published_year = ? WHERE id = ?"
	result, err := r.DB.Exec(query, book.Title, book.Author, book.Year, book.ID)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("no rows updated")
	}
	return nil
}

// DeleteBook deletes a book by its ID.
func (r *BookRepository) DeleteBook(id int64) error {
	query := "DELETE FROM books WHERE id = ?"
	result, err := r.DB.Exec(query, id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("no rows deleted")
	}
	return nil
}

// ListBooks returns all books.
func (r *BookRepository) ListBooks() ([]*models.Book, error) {
	query := "SELECT id, title, author, published_year FROM books"
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []*models.Book
	for rows.Next() {
		book := &models.Book{}
		if err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year); err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil
}
