package data

import (
	"database/sql"
	"fmt"
	"log"
)

type Book struct {
	ID     int64  `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

type BookModel struct {
	DB *sql.DB
}

func (bookModel *BookModel) GetAll() ([]*Book, error) {
	taleOfTwoCities := &Book{
		ID:     1001,
		Title:  "A Tale of Two Cities",
		Author: "Charles Dickens",
	}

	theSelfishGene := &Book{
		ID:     1002,
		Title:  "The Selfish Gene",
		Author: "Richard Dawkins",
	}

	loveInTheTimeOfCholera := &Book{
		ID:     1003,
		Title:  "Love in the Time of Cholera",
		Author: "Gabriel Garcia Marquez",
	}

	return []*Book{taleOfTwoCities, theSelfishGene, loveInTheTimeOfCholera}, nil
}

func (bookModel *BookModel) GetById(id int64) (*Book, error) {

	log.Println("retrieving book with id", id)
	query := `
        select id, title, author from books where id = $1;
    `

	var book Book
	err := bookModel.DB.QueryRow(query, id).Scan(&book.ID, &book.Title, &book.Author)
	if err != nil {
		if err == sql.ErrNoRows {
			// Handle the case where no rows are returned
			log.Printf("No book found with id %d", id)
			return nil, nil // or return a custom error
		}
		log.Fatalf("Cannot retrieve book with id %d: %v", id, err)
	}

	return &book, fmt.Errorf("failed to find book with id %d", id)
}

func (bookModel *BookModel) Insert(book *Book) error {
	query := `
         INSERT INTO books (TITLE, AUTHOR)
         values ($1, $2)
         returning id
    `
	arg := []interface{}{book.Title, book.Author}
	return bookModel.DB.QueryRow(query, arg).Scan(&book.ID)
}
